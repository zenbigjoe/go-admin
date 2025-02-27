package datamodel

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

// GetAllTypesTable return the model of table types.
func GetAllTypesTable() table.Table {

	allTypesTable := table.NewDefaultTable(nil, table.DefaultConfigWithDriver("mysql"))

	info := allTypesTable.GetInfo()

	info.AddField("Id", "id", db.Int)
	info.AddField("Type_1", "type_1", db.Tinyint)
	info.AddField("Type_2", "type_2", db.Smallint)
	info.AddField("Type_3", "type_3", db.Mediumint)
	info.AddField("Type_4", "type_4", db.Bigint)
	info.AddField("Type_5", "type_5", db.Float)
	info.AddField("Type_6", "type_6", db.Double)
	info.AddField("Type_7", "type_7", db.Double)
	info.AddField("Type_8", "type_8", db.Double)
	info.AddField("Type_9", "type_9", db.Decimal)
	info.AddField("Type_10", "type_10", db.Bit)
	info.AddField("Type_11", "type_11", db.Tinyint)
	info.AddField("Type_12", "type_12", db.Tinyint)
	info.AddField("Type_13", "type_13", db.Decimal)
	info.AddField("Type_14", "type_14", db.Decimal)
	info.AddField("Type_15", "type_15", db.Decimal)
	info.AddField("Type_16", "type_16", db.Char)
	info.AddField("Type_17", "type_17", db.Varchar)
	info.AddField("Type_18", "type_18", db.Tinytext)
	info.AddField("Type_19", "type_19", db.Text)
	info.AddField("Type_20", "type_20", db.Mediumtext)
	info.AddField("Type_21", "type_21", db.Longtext)
	info.AddField("Type_22", "type_22", db.Tinyblob)
	info.AddField("Type_23", "type_23", db.Mediumblob)
	info.AddField("Type_24", "type_24", db.Blob)
	info.AddField("Type_25", "type_25", db.Longblob)
	info.AddField("Type_26", "type_26", db.Binary)
	info.AddField("Type_27", "type_27", db.Varbinary)
	info.AddField("Type_28", "type_28", db.Enum)
	info.AddField("Type_29", "type_29", db.Set)
	info.AddField("Type_30", "type_30", db.Date)
	info.AddField("Type_31", "type_31", db.Datetime)
	info.AddField("Type_32", "type_32", db.Timestamp)
	info.AddField("Type_33", "type_33", db.Time)
	info.AddField("Type_34", "type_34", db.Year)
	info.AddField("Type_35", "type_35", db.Geometry)
	info.AddField("Type_36", "type_36", db.Point)
	info.AddField("Type_39", "type_39", db.Multilinestring)
	info.AddField("Type_41", "type_41", db.Multipolygon)
	info.AddField("Type_37", "type_37", db.Linestring)
	info.AddField("Type_38", "type_38", db.Polygon)
	info.AddField("Type_40", "type_40", db.Multipoint)
	info.AddField("Type_42", "type_42", db.Geometrycollection)
	info.AddField("Type_50", "type_50", db.Double)
	info.AddField("Type_51", "type_51", db.JSON)

	info.SetTable("all_types").SetTitle("All_types").SetDescription("All_types")

	formList := allTypesTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Type_1", "type_1", db.Tinyint, form.Number)
	formList.AddField("Type_2", "type_2", db.Smallint, form.Number)
	formList.AddField("Type_3", "type_3", db.Mediumint, form.Number)
	formList.AddField("Type_4", "type_4", db.Bigint, form.Number)
	formList.AddField("Type_5", "type_5", db.Float, form.Text)
	formList.AddField("Type_6", "type_6", db.Double, form.Text)
	formList.AddField("Type_7", "type_7", db.Double, form.Text)
	formList.AddField("Type_8", "type_8", db.Double, form.Text)
	formList.AddField("Type_9", "type_9", db.Decimal, form.Text)
	formList.AddField("Type_10", "type_10", db.Bit, form.Text)
	formList.AddField("Type_11", "type_11", db.Tinyint, form.Number)
	formList.AddField("Type_12", "type_12", db.Tinyint, form.Number)
	formList.AddField("Type_13", "type_13", db.Decimal, form.Text)
	formList.AddField("Type_14", "type_14", db.Decimal, form.Text)
	formList.AddField("Type_15", "type_15", db.Decimal, form.Text)
	formList.AddField("Type_16", "type_16", db.Char, form.Text)
	formList.AddField("Type_17", "type_17", db.Varchar, form.Text)
	formList.AddField("Type_18", "type_18", db.Tinytext, form.RichText)
	formList.AddField("Type_19", "type_19", db.Text, form.RichText)
	formList.AddField("Type_20", "type_20", db.Mediumtext, form.RichText)
	formList.AddField("Type_21", "type_21", db.Longtext, form.RichText)
	formList.AddField("Type_22", "type_22", db.Tinyblob, form.Text)
	formList.AddField("Type_23", "type_23", db.Mediumblob, form.Text)
	formList.AddField("Type_24", "type_24", db.Blob, form.Text)
	formList.AddField("Type_25", "type_25", db.Longblob, form.Text)
	formList.AddField("Type_26", "type_26", db.Binary, form.Text)
	formList.AddField("Type_27", "type_27", db.Varbinary, form.Text)
	formList.AddField("Type_28", "type_28", db.Enum, form.Text)
	formList.AddField("Type_29", "type_29", db.Set, form.Text)
	formList.AddField("Type_30", "type_30", db.Date, form.Datetime)
	formList.AddField("Type_31", "type_31", db.Datetime, form.Datetime)
	formList.AddField("Type_32", "type_32", db.Timestamp, form.Datetime)
	formList.AddField("Type_33", "type_33", db.Time, form.Datetime)
	formList.AddField("Type_34", "type_34", db.Year, form.Datetime)
	formList.AddField("Type_35", "type_35", db.Geometry, form.Text)
	formList.AddField("Type_36", "type_36", db.Point, form.Text)
	formList.AddField("Type_39", "type_39", db.Multilinestring, form.Text)
	formList.AddField("Type_41", "type_41", db.Multipolygon, form.Text)
	formList.AddField("Type_37", "type_37", db.Linestring, form.Text)
	formList.AddField("Type_38", "type_38", db.Polygon, form.Text)
	formList.AddField("Type_40", "type_40", db.Multipoint, form.Text)
	formList.AddField("Type_42", "type_42", db.Geometrycollection, form.Text)
	formList.AddField("Type_50", "type_50", db.Double, form.Text)
	formList.AddField("Type_51", "type_51", db.JSON, form.Text)

	formList.SetTable("all_types").SetTitle("All_types").SetDescription("All_types")

	return allTypesTable
}
