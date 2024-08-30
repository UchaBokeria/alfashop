package migration

import "main/pkg/storage"

func FilterEnums() {
	storage.DB.Exec(`CREATE TYPE filter_types AS ENUM (
		'range',
		'checkbox',
		'select',
		'radio',
		'text',
		'date',
		'number',
		'color',
		'switch',
		'multiselect',
		'rating',
		'search'
	);`)
}
