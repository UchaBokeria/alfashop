package categories

import (
	"main/internal/model"
)

var Automobile = []model.Category_filters{
	{
		Name:             "Choose a Range",
		Slug:             "choose-a-range",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "alfashop Lube", Key: "alfashop-lube", Value: "alfashop Lube"},
			{Name: "Galaxy Compettion", Key: "galaxy-competition", Value: "Galaxy Compettion"},
			{Name: "Collection Lubes", Key: "collection-lubes", Value: "Collection Lubes"},
			{Name: "alfashop VX Premium", Key: "alfashop-vx-premium", Value: "alfashop VX Premium"},
			{Name: "alfashop VX", Key: "alfashop-vx", Value: "alfashop VX"},
			{Name: "alfashop Pro", Key: "alfashop-pro", Value: "alfashop Pro"},
		},
	},
	{
		Name:             "2Page",
		Slug:             "2page",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "Synthetic 100%", Key: "synthetic-100", Value: "Synthetic 100%"},
			{Name: "Mineral", Key: "mineral", Value: "Mineral"},
			{Name: "Synthetic", Key: "synthetic", Value: "Synthetic"},
			{Name: "Synthetic Technology", Key: "synthetic-technology", Value: "Synthetic Technology"},
			{Name: "Semi Synthetic", Key: "semi-synthetic", Value: "Semi Synthetic"},
		},
	},
	{
		Name:             "Viscosity",
		Slug:             "viscosity",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "30", Key: "30", Value: "30"},
			{Name: "5W30", Key: "5w30", Value: "5W30"},
			{Name: "0W40", Key: "0w40", Value: "0W40"},
			{Name: "0W16", Key: "0w16", Value: "0W16"},
			{Name: "10W60", Key: "10w60", Value: "10W60"},
			{Name: "5W20", Key: "5w20", Value: "5W20"},
			{Name: "0W8", Key: "0w8", Value: "0W8"},
			{Name: "0W30", Key: "0w30", Value: "0W30"},
			{Name: "0W20", Key: "0w20", Value: "0W20"},
			{Name: "5W40", Key: "5w40", Value: "5W40"},
			{Name: "10W50", Key: "10w50", Value: "10W50"},
			{Name: "5W50", Key: "5w50", Value: "5W50"},
			{Name: "15W50", Key: "15w50", Value: "15W50"},
			{Name: "10W40", Key: "10w40", Value: "10W40"},
			{Name: "20W50", Key: "20w50", Value: "20W50"},
			{Name: "15W40", Key: "15w40", Value: "15W40"},
			{Name: "10W30", Key: "10w30", Value: "10W30"},
		},
	},
	{
		Name:             "Engine",
		Slug:             "engine",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "4-Stoke", Key: "4-stoke", Value: "4-Stoke"},
			{Name: "2-Stoke", Key: "2-stoke", Value: "2-Stoke"},
		},
	},
	{
		Name:             "Packaging",
		Slug:             "packaging",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "60L Barrel", Key: "60l-barrel", Value: "60L Barrel"},
			{Name: "208L Barrel", Key: "208l-barrel", Value: "208L Barrel"},
			{Name: "2L Can", Key: "2l-can", Value: "2L Can"},
			{Name: "1000L Tank", Key: "1000l-tank", Value: "1000L Tank"},
			{Name: "1000L Bulk", Key: "1000l-bulk", Value: "1000L Bulk"},
			{Name: "1L Can", Key: "1l-can", Value: "1L Can"},
			{Name: "4L Can", Key: "4l-can", Value: "4L Can"},
			{Name: "5L Can", Key: "5l-can", Value: "5L Can"},
			{Name: "250ML Small Can", Key: "250ml-small-can", Value: "250ML Small Can"},
			{Name: "20L Barrelrel", Key: "20l-barrelrel", Value: "20L Barrelrel"},
			{Name: "Bag in Box 20L", Key: "bag-in-box-20l", Value: "Bag in Box 20L"},
		},
	},
}

var MotoQuadKarting = []model.Category_filters{
	{
		Name:             "Choose a Range",
		Slug:             "choose-a-range",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "4 Stoke Motorcycle Lubes", Key: "4-stoke-motorcycle-lubes", Value: "4 Stoke Motorcycle Lubes"},
			{Name: "2 Stoke Motorcycle Lubes", Key: "2-stoke-motorcycle-lubes", Value: "2 Stoke Motorcycle Lubes"},
			{Name: "Scooter Lubes", Key: "scooter-lubes", Value: "Scooter Lubes"},
			{Name: "Quad Lubes", Key: "quad-lubes", Value: "Quad Lubes"},
			{Name: "Karting Lubes", Key: "karting-lubes", Value: "Karting Lubes"},
			{Name: "Snowmobile Lubes", Key: "snowmobile-lubes", Value: "Snowmobile Lubes"},
			{Name: "Motorcycles Specific Product", Key: "motorcycles-specific-product", Value: "Motorcycles Specific Product"},
			{Name: "Motorcyclesgearbox Lubes", Key: "motorcyclesgearbox-lubes", Value: "Motorcyclesgearbox Lubes"},
			{Name: "Motorcycles Fork Lubes", Key: "motorcycles-fork-lubes", Value: "Motorcycles Fork Lubes"},
		},
	},
	{
		Name:             "2Page",
		Slug:             "2page",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "Synthetic 100%", Key: "synthetic-100", Value: "Synthetic 100%"},
			{Name: "Mineral", Key: "mineral", Value: "Mineral"},
			{Name: "Synthetic", Key: "synthetic", Value: "Synthetic"},
			{Name: "Synthetic Technology", Key: "synthetic-technology", Value: "Synthetic Technology"},
			{Name: "Semi Synthetic", Key: "semi-synthetic", Value: "Semi Synthetic"},
		},
	},
	{
		Name:             "Engine",
		Slug:             "engine",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "4-Stoke", Key: "4-stoke", Value: "4-Stoke"},
			{Name: "2-Stoke", Key: "2-stoke", Value: "2-Stoke"},
		},
	},
	{
		Name:             "Viscosity",
		Slug:             "viscosity",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "10W40", Key: "10w40", Value: "10W40"},
			{Name: "10W30", Key: "10w30", Value: "10W30"},
			{Name: "10W60", Key: "10w60", Value: "10W60"},
			{Name: "15W50", Key: "15w50", Value: "15W50"},
			{Name: "10W50", Key: "10w50", Value: "10W50"},
			{Name: "5W40", Key: "5w40", Value: "5W40"},
			{Name: "20W50", Key: "20w50", Value: "20W50"},
			{Name: "80W90", Key: "80w90", Value: "80W90"},
			{Name: "0W40", Key: "0w40", Value: "0W40"},
		},
	},
	{
		Name:             "Packaging",
		Slug:             "packaging",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "1L Can", Key: "1l-can", Value: "1L Can"},
			{Name: "208L Barrel", Key: "208l-barrel", Value: "208L Barrel"},
			{Name: "4L Can", Key: "4l-can", Value: "4L Can"},
			{Name: "2L Can", Key: "2l-can", Value: "2L Can"},
			{Name: "60L Barrel", Key: "60l-barrel", Value: "60L Barrel"},
			{Name: "1000L Bulk", Key: "1000l-bulk", Value: "1000L Bulk"},
			{Name: "400ML Spray", Key: "400ml-spray", Value: "400ML Spray"},
			{Name: "5L Can", Key: "5l-can", Value: "5L Can"},
			{Name: "125ML Tube", Key: "125ml-tube", Value: "125ML Tube"},
			{Name: "500ML Spray", Key: "500ml-spray", Value: "500ML Spray"},
		},
	},
}

var TransportHeavyEquipment = []model.Category_filters{
	{
		Name:             "Choose a Range",
		Slug:             "choose-a-range",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "Premium Mineral Lubes", Key: "premium-mineral-lubes", Value: "Premium Mineral Lubes"},
			{Name: "Hydraulic Systems", Key: "hydraulic-systems", Value: "Hydraulic Systems"},
			{Name: "Automatic Transmission Lubes", Key: "automatic-transmission-lubes", Value: "Automatic Transmission Lubes"},
			{Name: "Gear Boxes and Beam Axles", Key: "gear-boxes-and-beam-axles", Value: "Gear Boxes and Beam Axles"},
			{Name: "Brake and Suspension Systems", Key: "brake-and-suspension-systems", Value: "Brake and Suspension Systems"},
			{Name: "Cooling Systems", Key: "cooling-systems", Value: "Cooling Systems"},
		},
	},
	{
		Name:             "2Page",
		Slug:             "2page",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "Synthetic 100%", Key: "synthetic-100", Value: "Synthetic 100%"},
			{Name: "Mineral", Key: "mineral", Value: "Mineral"},
			{Name: "Synthetic", Key: "synthetic", Value: "Synthetic"},
			{Name: "Synthetic Technology", Key: "synthetic-technology", Value: "Synthetic Technology"},
			{Name: "Semi Synthetic", Key: "semi-synthetic", Value: "Semi Synthetic"},
		},
	},
	{
		Name:             "Viscosity",
		Slug:             "viscosity",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "10", Key: "10", Value: "10"},
			{Name: "30", Key: "30", Value: "30"},
			{Name: "40", Key: "40", Value: "40"},
			{Name: "50", Key: "50", Value: "50"},
			{Name: "80W90", Key: "80w90", Value: "80W90"},
			{Name: "5W30", Key: "5w30", Value: "5W30"},
			{Name: "10W40", Key: "10w40", Value: "10W40"},
			{Name: "15W40", Key: "15w40", Value: "15W40"},
			{Name: "10W30", Key: "10w30", Value: "10W30"},
			{Name: "20W50", Key: "20w50", Value: "20W50"},
		},
	},
	{
		Name:             "Engine",
		Slug:             "engine",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "4Stoke", Key: "4stoke", Value: "4Stoke"},
		},
	},
	{
		Name:             "Packaging",
		Slug:             "packaging",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "5L Can", Key: "5l-can", Value: "5L Can"},
			{Name: "1000L Tank", Key: "1000l-tank", Value: "1000L Tank"},
			{Name: "220K Barrel", Key: "220k-barrel", Value: "220K Barrel"},
			{Name: "60L Barrel", Key: "60l-barrel", Value: "60L Barrel"},
			{Name: "208L Barrel", Key: "208l-barrel", Value: "208L Barrel"},
			{Name: "Tonnelet 20L", Key: "tonnelet-20l", Value: "Tonnelet 20L"},
			{Name: "20L Barrel", Key: "20l-barrel", Value: "20L Barrel"},
			{Name: "2L Can", Key: "2l-can", Value: "2L Can"},
			{Name: "1L Can", Key: "1l-can", Value: "1L Can"},
			{Name: "500ML Small Can", Key: "500ml-small-can", Value: "500ML Small Can"},
		},
	},
}

var Farming = []model.Category_filters{
	{
		Name:             "Choose a Range",
		Slug:             "choose-a-range",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "Multifunctional Lubes", Key: "multifunctional-lubes", Value: "Multifunctional Lubes"},
			{Name: "Transmission Hydraulic Lubes", Key: "transmission-hydraulic-lubes", Value: "Transmission Hydraulic Lubes"},
			{Name: "2-Stoke Farming Equipment Lubes", Key: "2-stoke-farming-equipment-lubes", Value: "2-Stoke Farming Equipment Lubes"},
			{Name: "4-Stoke Farming Equipment Lubes", Key: "4-stoke-farming-equipment-lubes", Value: "4-Stoke Farming Equipment Lubes"},
			{Name: "Chainsaw Lubes", Key: "chainsaw-lubes", Value: "Chainsaw Lubes"},
			{Name: "Engine Lubes", Key: "engine-lubes", Value: "Engine Lubes"},
		},
	},
	{
		Name:             "2Page",
		Slug:             "2page",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "Synthetic 100%", Key: "synthetic-100", Value: "Synthetic 100%"},
			{Name: "Mineral", Key: "mineral", Value: "Mineral"},
			{Name: "Synthetic", Key: "synthetic", Value: "Synthetic"},
			{Name: "Synthetic Technology", Key: "synthetic-technology", Value: "Synthetic Technology"},
			{Name: "Semi Synthetic", Key: "semi-synthetic", Value: "Semi Synthetic"},
		},
	},
	{
		Name:             "Viscosity",
		Slug:             "viscosity",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "10W30", Key: "10w30", Value: "10W30"},
			{Name: "15W40", Key: "15w40", Value: "15W40"},
			{Name: "10W40", Key: "10w40", Value: "10W40"},
			{Name: "VG150", Key: "vg150", Value: "VG150"},
			{Name: "VG220", Key: "vg220", Value: "VG220"},
			{Name: "VG320", Key: "vg320", Value: "VG320"},
		},
	},
	{
		Name:             "Engine",
		Slug:             "engine",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "4Stoke", Key: "4stoke", Value: "4Stoke"},
			{Name: "2Stoke", Key: "2stoke", Value: "2Stoke"},
		},
	},
	{
		Name:             "Packaging",
		Slug:             "packaging",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "20L Barrel", Key: "20l-barrel", Value: "20L Barrel"},
			{Name: "60L Barrel", Key: "60l-barrel", Value: "60L Barrel"},
			{Name: "208L Barrel", Key: "208l-barrel", Value: "208L Barrel"},
			{Name: "5L Can", Key: "5l-can", Value: "5L Can"},
			{Name: "1000L Tank", Key: "1000l-tank", Value: "1000L Tank"},
			{Name: "1000L Bulk", Key: "1000l-bulk", Value: "1000L Bulk"},
			{Name: "2L Can", Key: "2l-can", Value: "2L Can"},
		},
	},
}

var GearboxesBeamAxles = []model.Category_filters{
	{
		Name:             "Choose a Range",
		Slug:             "choose-a-range",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "Automatic Transmission Lubes", Key: "automatic-transmission-lubes", Value: "Automatic Transmission Lubes"},
			{Name: "Synthetic Lubes", Key: "synthetic-lubes", Value: "Synthetic Lubes"},
			{Name: "Semi-Synthetic Lubes", Key: "semi-synthetic-lubes", Value: "Semi-Synthetic Lubes"},
			{Name: "Mineral Lubes", Key: "mineral-lubes", Value: "Mineral Lubes"},
			{Name: "Automatic Transmission Fluids", Key: "automatic-transmission-fluids", Value: "Automatic Transmission Fluids"},
			{Name: "Gearboxes Lubes", Key: "gearboxes-lubes", Value: "Gearboxes Lubes"},
		},
	},
	{
		Name:             "2Page",
		Slug:             "2page",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "Synthetic 100%", Key: "synthetic-100", Value: "Synthetic 100%"},
			{Name: "Mineral", Key: "mineral", Value: "Mineral"},
			{Name: "Synthetic", Key: "synthetic", Value: "Synthetic"},
			{Name: "Semi Synthetic", Key: "semi-synthetic", Value: "Semi Synthetic"},
		},
	},
	{
		Name:             "Viscosity",
		Slug:             "viscosity",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "90", Key: "90", Value: "90"},
			{Name: "220", Key: "220", Value: "220"},
			{Name: "320", Key: "320", Value: "320"},
			{Name: "75W", Key: "75w", Value: "75W"},
			{Name: "75W140", Key: "75w140", Value: "75W140"},
			{Name: "75W80", Key: "75w80", Value: "75W80"},
			{Name: "75W90", Key: "75w90", Value: "75W90"},
			{Name: "80W90", Key: "80w90", Value: "80W90"},
			{Name: "80W85", Key: "80w85", Value: "80W85"},
		},
	},
	{
		Name:             "Engine",
		Slug:             "engine",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "4-Stoke", Key: "4-stoke", Value: "4-Stoke"},
		},
	},
	{
		Name:             "Packaging",
		Slug:             "packaging",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "2L Can", Key: "2l-can", Value: "2L Can"},
			{Name: "60L Barrel", Key: "60l-barrel", Value: "60L Barrel"},
			{Name: "1L Can", Key: "1l-can", Value: "1L Can"},
			{Name: "208L Barrel", Key: "208l-barrel", Value: "208L Barrel"},
			{Name: "20L Barrel", Key: "20l-barrel", Value: "20L Barrel"},
			{Name: "5L Can", Key: "5l-can", Value: "5L Can"},
			{Name: "1000L Tank", Key: "1000l-tank", Value: "1000L Tank"},
		},
	},
}

var UpkeepAndCleaning = []model.Category_filters{
	{
		Name:             "Choose a Range",
		Slug:             "choose-a-range",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "Upkeep", Key: "upkeep", Value: "Upkeep"},
			{Name: "Protection", Key: "protection", Value: "Protection"},
			{Name: "Cleanser", Key: "cleanser", Value: "Cleanser"},
			{Name: "Grease", Key: "grease", Value: "Grease"},
		},
	},
	{
		Name:             "2Page",
		Slug:             "2page",
		Default_value_id: nil,
		Options:          []model.Category_filters_option{},
	},
	{
		Name:             "Viscosity",
		Slug:             "viscosity",
		Default_value_id: nil,
		Options:          []model.Category_filters_option{},
	},
	{
		Name:             "Engine",
		Slug:             "engine",
		Default_value_id: nil,
		Options:          []model.Category_filters_option{},
	},
	{
		Name:             "Packaging",
		Slug:             "packaging",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "500ML Spray 6A", Key: "500ml-spray-6a", Value: "500ML Spray 6A"},
			{Name: "Bldonnet 150ML", Key: "bldonnet-150ml", Value: "Bldonnet 150ML"},
			{Name: "250ML Small Can", Key: "250ml-small-can", Value: "250ML Small Can"},
			{Name: "400G Multipack", Key: "400g-multipack", Value: "400G Multipack"},
			{Name: "50K Gek", Key: "50k-gek", Value: "50K Gek"},
			{Name: "70 Bucket", Key: "70-bucket", Value: "70 Bucket"},
			{Name: "5L Can", Key: "5l-can", Value: "5L Can"},
			{Name: "25K Gek", Key: "25k-gek", Value: "25K Gek"},
			{Name: "100G Tube", Key: "100g-tube", Value: "100G Tube"},
			{Name: "125ML Spray", Key: "125ml-spray", Value: "125ML Spray"},
			{Name: "2.5L Can", Key: "2.5l-can", Value: "2.5L Can"},
			{Name: "Seau 4K", Key: "seau-4k", Value: "Seau 4K"},
			{Name: "Seau 9K", Key: "seau-9k", Value: "Seau 9K"},
			{Name: "400ML Spray", Key: "400ml-spray", Value: "400ML Spray"},
			{Name: "100ML Tube", Key: "100ml-tube", Value: "100ML Tube"},
			{Name: "100ML Small Can", Key: "100ml-small-can", Value: "100ML Small Can"},
			{Name: "Spray 500ML 12A", Key: "spray-500ml-12a", Value: "Spray 500ML 12A"},
			{Name: "20L Barrel", Key: "20l-barrel", Value: "20L Barrel"},
			{Name: "210L Barrel", Key: "210l-barrel", Value: "210L Barrel"},
			{Name: "150ML Spray", Key: "150ml-spray", Value: "150ML Spray"},
			{Name: "180K Barrel", Key: "180k-barrel", Value: "180K Barrel"},
			{Name: "1K Box", Key: "1k-box", Value: "1K Box"},
			{Name: "10K Bucket", Key: "10k-bucket", Value: "10K Bucket"},
			{Name: "50K Barrel", Key: "50k-barrel", Value: "50K Barrel"},
			{Name: "5K Bucket", Key: "5k-bucket", Value: "5K Bucket"},
			{Name: "15K Bucket", Key: "15k-bucket", Value: "15K Bucket"},
			{Name: "60L Barrel", Key: "60l-barrel", Value: "60L Barrel"},
			{Name: "208L Barrel", Key: "208l-barrel", Value: "208L Barrel"},
			{Name: "400 ML Small Can", Key: "400ml-small-can", Value: "400 ML Small Can"},
		},
	},
}

var SailingAndYachting = []model.Category_filters{
	{
		Name:             "Choose a Range",
		Slug:             "choose-a-range",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "2-stoke engine lubes", Key: "2-stoke-engine-lubes", Value: "2-stoke engine lubes"},
			{Name: "Gearboxes lubes", Key: "gearboxes-lubes", Value: "Gearboxes lubes"},
			{Name: "Special products", Key: "special-products", Value: "Special products"},
			{Name: "Transmission fluids", Key: "transmission-fluids", Value: "Transmission fluids"},
			{Name: "Hydraulic fluids", Key: "hydraulic-fluids", Value: "Hydraulic fluids"},
			{Name: "Greases", Key: "greases", Value: "Greases"},
			{Name: "4-stoke engine lubes", Key: "4-stoke-engine-lubes", Value: "4-stoke engine lubes"},
		},
	},
	{
		Name:             "2Page",
		Slug:             "2page",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "Synthetic 100%", Key: "synthetic-100", Value: "Synthetic 100%"},
			{Name: "Mineral", Key: "mineral", Value: "Mineral"},
			{Name: "Synthetic", Key: "synthetic", Value: "Synthetic"},
			{Name: "Semi synthetic", Key: "semi-synthetic", Value: "Semi synthetic"},
		},
	},
	{
		Name:             "Viscosity",
		Slug:             "viscosity",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "30", Key: "30", Value: "30"},
			{Name: "40", Key: "40", Value: "40"},
			{Name: "90", Key: "90", Value: "90"},
			{Name: "10W30", Key: "10w30", Value: "10W30"},
			{Name: "10W40", Key: "10w40", Value: "10W40"},
			{Name: "25W40", Key: "25w40", Value: "25W40"},
			{Name: "75W90", Key: "75w90", Value: "75W90"},
			{Name: "75W140", Key: "75w140", Value: "75W140"},
			{Name: "15W40", Key: "15w40", Value: "15W40"},
			{Name: "5W40", Key: "5w40", Value: "5W40"},
			{Name: "10W60", Key: "10w60", Value: "10W60"},
		},
	},
	{
		Name:             "Engine",
		Slug:             "engine",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "2-stoke", Key: "2-stoke", Value: "2-stoke"},
			{Name: "4-stoke", Key: "4-stoke", Value: "4-stoke"},
		},
	},
	{
		Name:             "Packaging",
		Slug:             "packaging",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "5L Can", Key: "5l-can", Value: "5L Can"},
			{Name: "1000L Tank", Key: "1000l-tank", Value: "1000L Tank"},
			{Name: "220K Barrel", Key: "220k-barrel", Value: "220K Barrel"},
			{Name: "2L Can", Key: "2l-can", Value: "2L Can"},
			{Name: "1L Can", Key: "1l-can", Value: "1L Can"},
			{Name: "60L Barrel", Key: "60l-barrel", Value: "60L Barrel"},
			{Name: "208L Barrel", Key: "208l-barrel", Value: "208L Barrel"},
			{Name: "200ML Small Can", Key: "200ml-small-can", Value: "200ML Small Can"},
			{Name: "20L Barrel", Key: "20l-barrel", Value: "20L Barrel"},
			{Name: "250ML Lube", Key: "250ml-lube", Value: "250ML Lube"},
			{Name: "1000L Bulk", Key: "1000l-bulk", Value: "1000L Bulk"},
			{Name: "400G Multipack", Key: "400g-multipack", Value: "400G Multipack"},
			{Name: "1K Box", Key: "1k-box", Value: "1K Box"},
			{Name: "180K Barrel", Key: "180k-barrel", Value: "180K Barrel"},
			{Name: "5K Bucket", Key: "5k-bucket", Value: "5K Bucket"},
			{Name: "50K Keg", Key: "50k-keg", Value: "50K Keg"},
		},
	},
}

var LightPlanes = []model.Category_filters{
	{
		Name:             "Choose a Range",
		Slug:             "choose-a-range",
		Default_value_id: nil,
		Options:          []model.Category_filters_option{},
	},
	{
		Name:             "2Page",
		Slug:             "2page",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "Synthetic 100%", Key: "synthetic-100", Value: "Synthetic 100%"},
			{Name: "Mineral", Key: "mineral", Value: "Mineral"},
			{Name: "Synthetic", Key: "synthetic", Value: "Synthetic"},
			{Name: "Semi synthetic", Key: "semi-synthetic", Value: "Semi synthetic"},
		},
	},
	{
		Name:             "Viscosity",
		Slug:             "viscosity",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "80", Key: "80", Value: "80"},
			{Name: "10W80", Key: "10w80", Value: "10W80"},
		},
	},
	{
		Name:             "Engine",
		Slug:             "engine",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "2-stoke", Key: "2-stoke", Value: "2-stoke"},
			{Name: "4-stoke", Key: "4-stoke", Value: "4-stoke"},
		},
	},
	{
		Name:             "Packaging",
		Slug:             "packaging",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "2L Can", Key: "2l-can", Value: "2L Can"},
			{Name: "1L Can", Key: "1l-can", Value: "1L Can"},
			{Name: "60L Barrel", Key: "60l-barrel", Value: "60L Barrel"},
		},
	},
}

var Specialities = []model.Category_filters{
	{
		Name:             "Choose a Range",
		Slug:             "choose-a-range",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "Cooling fluids", Key: "cooling-fluids", Value: "Cooling fluids"},
			{Name: "Synthetic brake fluid", Key: "synthetic-brake-fluid", Value: "Synthetic brake fluid"},
			{Name: "Hydraulic fluid", Key: "hydraulic-fluid", Value: "Hydraulic fluid"},
			{Name: "Windshield washer", Key: "windshield-washer", Value: "Windshield washer"},
			{Name: "Transmission lubes", Key: "transmission-lubes", Value: "Transmission lubes"},
			{Name: "Metalworking fluids", Key: "metalworking-fluids", Value: "Metalworking fluids"},
			{Name: "Coolant", Key: "coolant", Value: "Coolant"},
			{Name: "Gearboxes lubes", Key: "gearboxes-lubes", Value: "Gearboxes lubes"},
		},
	},
	{
		Name:             "2Page",
		Slug:             "2page",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "Mineral", Key: "mineral", Value: "Mineral"},
			{Name: "Synthetic", Key: "synthetic", Value: "Synthetic"},
			{Name: "Semi synthetic", Key: "semi-synthetic", Value: "Semi synthetic"},
		},
	},
	{
		Name:             "Viscosity",
		Slug:             "viscosity",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "220", Key: "220", Value: "220"},
			{Name: "320", Key: "320", Value: "320"},
		},
	},
	{
		Name:             "Engine",
		Slug:             "engine",
		Default_value_id: nil,
		Options:          []model.Category_filters_option{},
	},
	{
		Name:             "Packaging",
		Slug:             "packaging",
		Default_value_id: nil,
		Options: []model.Category_filters_option{
			{Name: "5L Can", Key: "5l-can", Value: "5L Can"},
			{Name: "1000L Tank", Key: "1000l-tank", Value: "1000L Tank"},
			{Name: "220K Barrel", Key: "220k-barrel", Value: "220K Barrel"},
			{Name: "500ML Small Can", Key: "500ml-small-can", Value: "500ML Small Can"},
			{Name: "1L Can", Key: "1l-can", Value: "1L Can"},
			{Name: "20L Barrel", Key: "20l-barrel", Value: "20L Barrel"},
			{Name: "210L Barrel", Key: "210l-barrel", Value: "210L Barrel"},
			{Name: "208L Barrel", Key: "208l-barrel", Value: "208L Barrel"},
			{Name: "60L Barrel", Key: "60l-barrel", Value: "60L Barrel"},
		},
	},
}
