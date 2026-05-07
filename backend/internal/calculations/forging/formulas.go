package forging

import "math"

type Result struct {
	ForgingForce     float64 `json:"forging_force"`
	ForgingPressure  float64 `json:"forging_pressure"`
	WorkDone         float64 `json:"work_done"`
	Power            float64 `json:"power"`
	DeformationSpeed float64 `json:"deformation_speed"`
	StrainDegree     float64 `json:"strain_degree"`
	InitialVolume    float64 `json:"initial_volume"`
	FinalVolume      float64 `json:"final_volume"`
	WorkpieceMass    float64 `json:"workpiece_mass"`
	ContactArea      float64 `json:"contact_area"`
	FinalDiameter    float64 `json:"final_diameter"`
	FinalSideA       float64 `json:"final_side_a"`
	FinalSideB       float64 `json:"final_side_b"`
	HeightReduction  float64 `json:"height_reduction"`
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

func Calculate(p Params) Result {
	props := getMaterialProps(p.Material, p.Temperature)

	var initialVolume, contactArea, finalA, finalB float64
	heightReduction := p.InitialHeight - p.FinalHeight

	if p.Shape == "cylinder" {
		initialVolume = math.Pi * p.DimensionA * p.DimensionA / 4.0 * p.InitialHeight
		finalA = math.Sqrt(4.0 * initialVolume / (math.Pi * p.FinalHeight))
		finalB = finalA
		contactArea = math.Pi * finalA * finalA / 4.0
	} else {
		initialVolume = p.DimensionA * p.DimensionB * p.InitialHeight
		finalB = p.DimensionB
		finalA = initialVolume / (finalB * p.FinalHeight)
		contactArea = finalA * finalB
	}

	var forgingPressure float64
	if p.Shape == "cylinder" {
		forgingPressure = props.YieldStress * (1 + p.FrictionCoeff*finalA/(3*p.FinalHeight))
	} else {
		eqDim := math.Sqrt(contactArea)
		forgingPressure = props.YieldStress * (1 + p.FrictionCoeff*eqDim/(2*p.FinalHeight))
	}

	forgingForce := forgingPressure * contactArea
	strainRate := p.DeformationSpeed / ((p.InitialHeight + p.FinalHeight) / 2.0)
	strainDegree := math.Log(p.InitialHeight / p.FinalHeight)

	shapeFactor := 1.5 + 0.5*p.FrictionCoeff
	workDone := forgingForce * heightReduction * shapeFactor / 1000.0

	deformationTime := heightReduction / p.DeformationSpeed
	if deformationTime < 0.001 {
		deformationTime = 0.001
	}
	power := workDone / deformationTime / 1000.0
	mass := props.Density * initialVolume / 1e6

	return Result{
		ForgingForce:     forgingForce,
		ForgingPressure:  forgingPressure,
		WorkDone:         workDone,
		Power:            power,
		DeformationSpeed: strainRate,
		StrainDegree:     strainDegree,
		InitialVolume:    initialVolume,
		FinalVolume:      initialVolume,
		WorkpieceMass:    mass,
		ContactArea:      contactArea,
		FinalDiameter:    finalA,
		FinalSideA:       finalA,
		FinalSideB:       finalB,
		HeightReduction:  heightReduction,
	}
}
