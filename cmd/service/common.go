package main

type sensor struct {
	CoretempIsa0000 struct {
		Adapter    string `json:"Adapter"`
		PackageId0 struct {
			Temp1Input     float64 `json:"temp1_input"`
			Temp1Max       float64 `json:"temp1_max"`
			Temp1Crit      float64 `json:"temp1_crit"`
			Temp1CritAlarm float64 `json:"temp1_crit_alarm"`
		} `json:"Package id 0"`
		Core0 struct {
			Temp2Input     float64 `json:"temp2_input"`
			Temp2Max       float64 `json:"temp2_max"`
			Temp2Crit      float64 `json:"temp2_crit"`
			Temp2CritAlarm float64 `json:"temp2_crit_alarm"`
		} `json:"Core 0"`
		Core1 struct {
			Temp3Input     float64 `json:"temp3_input"`
			Temp3Max       float64 `json:"temp3_max"`
			Temp3Crit      float64 `json:"temp3_crit"`
			Temp3CritAlarm float64 `json:"temp3_crit_alarm"`
		} `json:"Core 1"`
		Core2 struct {
			Temp4Input     float64 `json:"temp4_input"`
			Temp4Max       float64 `json:"temp4_max"`
			Temp4Crit      float64 `json:"temp4_crit"`
			Temp4CritAlarm float64 `json:"temp4_crit_alarm"`
		} `json:"Core 2"`
		Core3 struct {
			Temp5Input     float64 `json:"temp5_input"`
			Temp5Max       float64 `json:"temp5_max"`
			Temp5Crit      float64 `json:"temp5_crit"`
			Temp5CritAlarm float64 `json:"temp5_crit_alarm"`
		} `json:"Core 3"`
		Core4 struct {
			Temp6Input     float64 `json:"temp6_input"`
			Temp6Max       float64 `json:"temp6_max"`
			Temp6Crit      float64 `json:"temp6_crit"`
			Temp6CritAlarm float64 `json:"temp6_crit_alarm"`
		} `json:"Core 4"`
		Core5 struct {
			Temp7Input     float64 `json:"temp7_input"`
			Temp7Max       float64 `json:"temp7_max"`
			Temp7Crit      float64 `json:"temp7_crit"`
			Temp7CritAlarm float64 `json:"temp7_crit_alarm"`
		} `json:"Core 5"`
		Core6 struct {
			Temp8Input     float64 `json:"temp8_input"`
			Temp8Max       float64 `json:"temp8_max"`
			Temp8Crit      float64 `json:"temp8_crit"`
			Temp8CritAlarm float64 `json:"temp8_crit_alarm"`
		} `json:"Core 6"`
		Core7 struct {
			Temp9Input     float64 `json:"temp9_input"`
			Temp9Max       float64 `json:"temp9_max"`
			Temp9Crit      float64 `json:"temp9_crit"`
			Temp9CritAlarm float64 `json:"temp9_crit_alarm"`
		} `json:"Core 7"`
	} `json:"coretemp-isa-0000"`
	DellDdvVirtual0 struct {
		Adapter string `json:"Adapter"`
		CPUFan  struct {
			Fan1Input float64 `json:"fan1_input"`
		} `json:"CPU Fan"`
		VideoFan struct {
			Fan2Input float64 `json:"fan2_input"`
		} `json:"Video Fan"`
		CPU struct {
			Temp1Input float64 `json:"temp1_input"`
			Temp1Max   float64 `json:"temp1_max"`
			Temp1Min   float64 `json:"temp1_min"`
		} `json:"CPU"`
		SODIMM struct {
			Temp2Input float64 `json:"temp2_input"`
			Temp2Max   float64 `json:"temp2_max"`
			Temp2Min   float64 `json:"temp2_min"`
		} `json:"SODIMM"`
		Ambient struct {
			Temp3Input float64 `json:"temp3_input,omitempty"`
			Temp3Max   float64 `json:"temp3_max,omitempty"`
			Temp3Min   float64 `json:"temp3_min,omitempty"`
			Temp6Input float64 `json:"temp6_input,omitempty"`
			Temp6Max   float64 `json:"temp6_max,omitempty"`
			Temp6Min   float64 `json:"temp6_min,omitempty"`
			Temp8Input float64 `json:"temp8_input,omitempty"`
			Temp8Max   float64 `json:"temp8_max,omitempty"`
			Temp8Min   float64 `json:"temp8_min,omitempty"`
			Temp9Input float64 `json:"temp9_input,omitempty"`
			Temp9Max   float64 `json:"temp9_max,omitempty"`
			Temp9Min   float64 `json:"temp9_min,omitempty"`
		} `json:"Ambient"`
		Unknown struct {
			Temp4Input  float64 `json:"temp4_input,omitempty"`
			Temp4Max    float64 `json:"temp4_max,omitempty"`
			Temp4Min    float64 `json:"temp4_min,omitempty"`
			Temp7Input  float64 `json:"temp7_input,omitempty"`
			Temp7Max    float64 `json:"temp7_max,omitempty"`
			Temp7Min    float64 `json:"temp7_min,omitempty"`
			Temp11Input float64 `json:"temp11_input,omitempty"`
			Temp11Max   float64 `json:"temp11_max,omitempty"`
			Temp11Min   float64 `json:"temp11_min,omitempty"`
		} `json:"Unknown"`
		HDD struct {
			Temp5Input float64 `json:"temp5_input"`
			Temp5Max   float64 `json:"temp5_max"`
			Temp5Min   float64 `json:"temp5_min"`
		} `json:"HDD"`
		Other struct {
			Temp10Input float64 `json:"temp10_input"`
			Temp10Max   float64 `json:"temp10_max"`
			Temp10Min   float64 `json:"temp10_min"`
		} `json:"Other"`
		Video struct {
			Temp12Input float64 `json:"temp12_input"`
			Temp12Max   float64 `json:"temp12_max"`
			Temp12Min   float64 `json:"temp12_min"`
		} `json:"Video"`
	} `json:"dell_ddv-virtual-0"`
	UcsiSourcePsyUSBC000002Isa0000 struct {
		Adapter string `json:"Adapter"`
		In0     struct {
			In0Input float64 `json:"in0_input"`
			In0Min   float64 `json:"in0_min"`
			In0Max   float64 `json:"in0_max"`
		} `json:"in0"`
		Curr1 struct {
			Curr1Input float64 `json:"curr1_input"`
			Curr1Max   float64 `json:"curr1_max"`
		} `json:"curr1"`
	} `json:"ucsi_source_psy_USBC000:002-isa-0000"`
	NvmePciE100 struct {
		Adapter   string `json:"Adapter"`
		Composite struct {
			Temp1Input float64 `json:"temp1_input"`
			Temp1Max   float64 `json:"temp1_max"`
			Temp1Min   float64 `json:"temp1_min"`
			Temp1Crit  float64 `json:"temp1_crit"`
			Temp1Alarm float64 `json:"temp1_alarm"`
		} `json:"Composite"`
	} `json:"nvme-pci-e100"`
	Iwlwifi1Virtual0 struct {
		Adapter string `json:"Adapter"`
		Temp1   struct {
			Temp1Input float64 `json:"temp1_input"`
		} `json:"temp1"`
	} `json:"iwlwifi_1-virtual-0"`
	DellSmmVirtual0 struct {
		Adapter string `json:"Adapter"`
		Fan1    struct {
			Fan1Input float64 `json:"fan1_input"`
			Fan1Min   float64 `json:"fan1_min"`
			Fan1Max   float64 `json:"fan1_max"`
		} `json:"fan1"`
		Fan2 struct {
			Fan2Input float64 `json:"fan2_input"`
			Fan2Min   float64 `json:"fan2_min"`
			Fan2Max   float64 `json:"fan2_max"`
		} `json:"fan2"`
		Temp1 struct {
			Temp1Input float64 `json:"temp1_input"`
		} `json:"temp1"`
		Temp2 struct {
			Temp2Input float64 `json:"temp2_input"`
		} `json:"temp2"`
		Temp3 struct {
			Temp3Input float64 `json:"temp3_input"`
		} `json:"temp3"`
		Temp4 struct {
			Temp4Input float64 `json:"temp4_input"`
		} `json:"temp4"`
		Temp5 struct {
			Temp5Input float64 `json:"temp5_input"`
		} `json:"temp5"`
		Temp6 struct {
			Temp6Input float64 `json:"temp6_input"`
		} `json:"temp6"`
		Temp7 struct {
			Temp7Input float64 `json:"temp7_input"`
		} `json:"temp7"`
		Temp8 struct {
			Temp8Input float64 `json:"temp8_input"`
		} `json:"temp8"`
		Temp9 struct {
			Temp9Input float64 `json:"temp9_input"`
		} `json:"temp9"`
		Temp10 struct {
			Temp10Input float64 `json:"temp10_input"`
		} `json:"temp10"`
	} `json:"dell_smm-virtual-0"`
	UcsiSourcePsyUSBC000003Isa0000 struct {
		Adapter string `json:"Adapter"`
		In0     struct {
			In0Input float64 `json:"in0_input"`
			In0Min   float64 `json:"in0_min"`
			In0Max   float64 `json:"in0_max"`
		} `json:"in0"`
		Curr1 struct {
			Curr1Input float64 `json:"curr1_input"`
			Curr1Max   float64 `json:"curr1_max"`
		} `json:"curr1"`
	} `json:"ucsi_source_psy_USBC000:003-isa-0000"`
	UcsiSourcePsyUSBC000001Isa0000 struct {
		Adapter string `json:"Adapter"`
		In0     struct {
			In0Input float64 `json:"in0_input"`
			In0Min   float64 `json:"in0_min"`
			In0Max   float64 `json:"in0_max"`
		} `json:"in0"`
		Curr1 struct {
			Curr1Input float64 `json:"curr1_input"`
			Curr1Max   float64 `json:"curr1_max"`
		} `json:"curr1"`
	} `json:"ucsi_source_psy_USBC000:001-isa-0000"`
	BAT0Acpi0 struct {
		Adapter string `json:"Adapter"`
		In0     struct {
			In0Input float64 `json:"in0_input"`
		} `json:"in0"`
		Curr1 struct {
			Curr1Input float64 `json:"curr1_input"`
		} `json:"curr1"`
	} `json:"BAT0-acpi-0"`
}
