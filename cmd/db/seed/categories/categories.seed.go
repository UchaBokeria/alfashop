package categories

import (
	"main/internal/model"
	"main/pkg/storage"
)

func GetPointedUInt(num int) *uint {
	unum := uint(num)
	return &unum
}

var Seed = []model.Categories{
	// parents
	{
		IconID: 15,
		Public: true,
		Name: "ძრავის ზეთი",
		Slug: "motorenoel",
	},
	{
		IconID: 15,
		Public: true,
		Name: "ტრანსმისიური ზეთი",
		Slug: "getriebeoel",
	},
	{
		IconID: 15,
		Public: true,
		Name: "ჰიდრავლიკური ზეთი",
		Slug: "hydraulikoel",
	},

	// orphants-1
	{
		IconID: 15,
		Public: true,
		Name: "რადიატორული ანტიფრიზი",
		Slug: "kuehlerfrostschutz",
	},
	{
		IconID: 15,
		Public: true,
		Name: "დანამატები",
		Slug: "additive",
	},
	{
		IconID: 15,
		Public: true,
		Name: "სამუხრუჭე სითხე",
		Slug: "bremsflussigkeit",
	},
	{
		IconID: 15,
		Public: true,
		Name: "სატრანსპორტო საშუალების მოვლის დამლაგებელი",
		Slug: "fahrzeugpflege-reiniger",
	},
	{
		IconID: 15,
		Public: true,
		Name: "ზამთრის ქიმიური აქსესუარები",
		Slug: "winterchemie-zubehoer",
	},

	// orphants-2
	{
		IconID: 15,
		ParentID: GetPointedUInt(0),
		Public: true,
		Name: "სამრეწველო ზეთი",
		Slug: "industrieoel",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(0),
		Public: true,
		Name: "მსუქანი",
		Slug: "fette",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(0),
		Public: true,
		Name: "ყველა სარეკლამო ნივთები",
		Slug: "alle-werbeartikel",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(0),
		Public: true,
		Name: "აქსესუარები",
		Slug: "zubehoer",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(0),
		Public: true,
		Name: "სალბის ჯაჭვის ზეთი",
		Slug: "sagekettenol",
	},
	
	// motorenoel
	{
		IconID: 15,
		ParentID: GetPointedUInt(1),
		Public: true,
		Name: "კლასიკური ძრავის ზეთი",
		Slug: "classic-motorenoel",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(1),
		Public: true,
		Name: "2 ტაქტიანი ძრავის ზეთი",
		Slug: "2-takt-motorenoel",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(1),
		Public: true,
		Name: "4 ტაქტიანი ძრავის ზეთი",
		Slug: "4-takt-motorenoel",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(1),
		Public: true,
		Name: "საავტომობილო ზეთი სასოფლო-სამეურნეო-სატრანსპორტო სატრანსპორტო საშუალებებისა და სამშენებლო მანქანებისთვის",
		Slug: "motoroel-fuer-landwirtsch-fzg-und-baumaschinen",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(1),
		Public: true,
		Name: "საზღვაო ძრავის ზეთი",
		Slug: "marine-motoroel",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(1),
		Public: true,
		Name: "სხვა საავტომობილო ზეთი",
		Slug: "sonstiges-motorenoel",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(1),
		Public: true,
		Name: "მანქანის ძრავის ზეთი",
		Slug: "pkw-motorenoel",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(1),
		Public: true,
		Name: "სატვირთო მანქანის ძრავის ზეთი",
		Slug: "lkw-motorenoel",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(1),
		Public: true,
		Name: "მოტოციკლის ძრავის ზეთი",
		Slug: "motorrad-motorenoel",
	},

	// getriebeoel
	{
		IconID: 15,
		ParentID: GetPointedUInt(2),
		Public: true,
		Name: "გადაცემათა კოლოფი ავტომატური გადაცემისთვის",
		Slug: "getriebeoel-fur-automatikgetriebe",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(2),
		Public: true,
		Name: "გადაცემათა ზეთები მექანიკური გადაცემათა კოლოფისთვის და წამყვანი ღერძებისთვის",
		Slug: "getriebeoele-fur-schaltgetriebe-und-antriebsachsen",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(2),
		Public: true,
		Name: "საზღვაო გადაცემათა ზეთი",
		Slug: "marine-getriebeol",
	},

	// hydraulikoel
	{
		IconID: 15,
		ParentID: GetPointedUInt(3),
		Public: true,
		Name: "მანქანის ჰიდრავლიკური ზეთი",
		Slug: "pkw-hydraulikoel",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(3),
		Public: true,
		Name: "მოტოციკლის ჰიდრავლიკური ზეთი",
		Slug: "motorrad-hydraulikoel",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(3),
		Public: true,
		Name: "ჰიდრავლიკური ზეთები სასოფლო-სამეურნეო მანქანებისთვის და სამშენებლო მანქანებისთვის",
		Slug: "hydraulikole-fur-landw-fzg-und-baumaschinen",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(3),
		Public: true,
		Name: "საზღვაო ჰიდრავლიკური ზეთი",
		Slug: "marine-hydraulikoel",
	},
	{
		IconID: 15,
		ParentID: GetPointedUInt(3),
		Public: true,
		Name: "სხვა-ჰიდრავლიკური ზეთი",
		Slug: "sonstiges-hydraulikoel",
	},
}

func Populate() {
	for _, row := range Seed { storage.DB.Create(&row) }
}