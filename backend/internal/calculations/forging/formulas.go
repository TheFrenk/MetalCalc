package forging

import "math"

type PassResult struct {
	Pass        int     `json:"pass"`
	HeightStart float64 `json:"height_start"`
	HeightEnd   float64 `json:"height_end"`
	Force       float64 `json:"force"`
	Pressure    float64 `json:"pressure"`
	Work        float64 `json:"work"`
	ContactArea float64 `json:"contact_area"`
}

type Result struct {
	ForgingForce     float64      `json:"forging_force"`
	ForgingPressure  float64      `json:"forging_pressure"`
	WorkDone         float64      `json:"work_done"`
	Power            float64      `json:"power"`
	DeformationSpeed float64      `json:"deformation_speed"`
	StrainDegree     float64      `json:"strain_degree"`
	InitialVolume    float64      `json:"initial_volume"`
	FinalVolume      float64      `json:"final_volume"`
	WorkpieceMass    float64      `json:"workpiece_mass"`
	ContactArea      float64      `json:"contact_area"`
	FinalDiameter    float64      `json:"final_diameter"`
	FinalSideA       float64      `json:"final_side_a"`
	FinalSideB       float64      `json:"final_side_b"`
	HeightReduction  float64      `json:"height_reduction"`
	Passes           int          `json:"passes"`
	PassResults      []PassResult `json:"pass_results"`
}

type Params struct {
	Material         string  `json:"material"`
	Shape            string  `json:"shape"`
	DimensionA       float64 `json:"dimension_a"`
	DimensionB       float64 `json:"dimension_b"`
	InitialHeight    float64 `json:"initial_height"`
	FinalHeight      float64 `json:"final_height"`
	Temperature      float64 `json:"temperature"`
	FrictionCoeff    float64 `json:"friction_coeff"`
	DeformationSpeed float64 `json:"deformation_speed"`
	Passes           int     `json:"passes"`
}

type MaterialProps struct {
	YieldStress float64
	Density     float64
}

func getMaterialProps(material string, temp float64) MaterialProps {
	switch material {
	case "steel":
		sy := 600.0
		if temp > 900 {
			sy = 80.0
		} else if temp > 600 {
			sy = 200.0
		}
		return MaterialProps{YieldStress: sy, Density: 7.85}
	case "aluminum":
		sy := 250.0
		if temp > 350 {
			sy = 30.0
		} else if temp > 200 {
			sy = 80.0
		}
		return MaterialProps{YieldStress: sy, Density: 2.70}
	case "copper":
		sy := 300.0
		if temp > 700 {
			sy = 40.0
		} else if temp > 400 {
			sy = 120.0
		}
		return MaterialProps{YieldStress: sy, Density: 8.96}
	case "titanium":
		sy := 800.0
		if temp > 900 {
			sy = 150.0
		} else if temp > 600 {
			sy = 400.0
		}
		return MaterialProps{YieldStress: sy, Density: 4.51}
	default:
		return MaterialProps{YieldStress: 600.0, Density: 7.85}
	}
}

func calculatePass(shape string, mu, h0, h1, dimA, dimB float64, props MaterialProps, vSpeed float64) (force, pressure, work, power, contactArea, finalA, finalB float64) {
	volume := 0.0
	if shape == "cylinder" {
		volume = math.Pi * dimA * dimA / 4.0 * h0
		finalA = math.Sqrt(4.0 * volume / (math.Pi * h1))
		finalB = finalA
		contactArea = math.Pi * finalA * finalA / 4.0
		pressure = props.YieldStress * (1.0 + mu*finalA/(3.0*h1))
	} else {
		volume = dimA * dimB * h0
		finalB = dimB
		finalA = volume / (finalB * h1)
		contactArea = finalA * finalB
		minDim := math.Min(finalA, finalB)
		pressure = props.YieldStress * (1.0 + mu*minDim/(2.0*h1))
	}

	force = pressure * contactArea

	// Зусилля на початку проходу для розрахунку середньої роботи
	var f0 float64
	if shape == "cylinder" {
		initArea := math.Pi * dimA * dimA / 4.0
		p0 := props.YieldStress * (1.0 + mu*dimA/(3.0*h0))
		f0 = p0 * initArea
	} else {
		minDim0 := math.Min(dimA, dimB)
		p0 := props.YieldStress * (1.0 + mu*minDim0/(2.0*h0))
		f0 = p0 * (dimA * dimB)
	}

	deltaH := h0 - h1
	work = (f0 + force) / 2.0 * deltaH / 1e3 // Н·мм → Дж

	dt := deltaH / vSpeed
	if dt < 0.001 {
		dt = 0.001
	}
	power = work / dt / 1e3 // Дж → кВт

	return
}

func Calculate(p Params) Result {
	passes := p.Passes
	if passes < 1 {
		passes = 1
	}

	props := getMaterialProps(p.Material, p.Temperature)

	totalStrain := math.Log(p.InitialHeight / p.FinalHeight)
	strainPerPass := totalStrain / float64(passes)

	var (
		totalWork       float64
		totalForce      float64
		totalPressure   float64
		totalArea       float64
		totalPower      float64
		totalStrainRate float64
	)

	curH := p.InitialHeight
	curA := p.DimensionA
	curB := p.DimensionB
	var lastFinalA, lastFinalB float64

	passResults := make([]PassResult, 0, passes)

	for i := 0; i < passes; i++ {
		nextH := curH / math.Exp(strainPerPass)
		if i == passes-1 {
			nextH = p.FinalHeight
		}

		f, pr, w, pw, area, fa, fb := calculatePass(
			p.Shape, p.FrictionCoeff,
			curH, nextH,
			curA, curB,
			props, p.DeformationSpeed,
		)

		passResults = append(passResults, PassResult{
			Pass:        i + 1,
			HeightStart: curH,
			HeightEnd:   nextH,
			Force:       f,
			Pressure:    pr,
			Work:        w,
			ContactArea: area,
		})

		totalForce += f
		totalPressure += pr
		totalWork += w
		totalPower += pw
		totalArea += area
		totalStrainRate += p.DeformationSpeed / ((curH + nextH) / 2.0)

		lastFinalA = fa
		lastFinalB = fb

		curH = nextH
		curA = fa
		curB = fb
	}

	avgForce := totalForce / float64(passes)
	avgPressure := totalPressure / float64(passes)
	avgArea := totalArea / float64(passes)
	avgStrainRate := totalStrainRate / float64(passes)

	var initialVolume float64
	if p.Shape == "cylinder" {
		initialVolume = math.Pi * p.DimensionA * p.DimensionA / 4.0 * p.InitialHeight
	} else {
		initialVolume = p.DimensionA * p.DimensionB * p.InitialHeight
	}

	mass := props.Density * initialVolume / 1e6
	strainDegree := math.Log(p.InitialHeight / p.FinalHeight)

	return Result{
		ForgingForce:     avgForce,
		ForgingPressure:  avgPressure,
		WorkDone:         totalWork,
		Power:            totalPower,
		DeformationSpeed: avgStrainRate,
		StrainDegree:     strainDegree,
		InitialVolume:    initialVolume,
		FinalVolume:      initialVolume,
		WorkpieceMass:    mass,
		ContactArea:      avgArea,
		FinalDiameter:    lastFinalA,
		FinalSideA:       lastFinalA,
		FinalSideB:       lastFinalB,
		HeightReduction:  p.InitialHeight - p.FinalHeight,
		Passes:           passes,
		PassResults:      passResults,
	}
}
