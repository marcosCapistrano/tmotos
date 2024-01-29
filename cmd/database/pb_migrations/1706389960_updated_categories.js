/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("r76mw51ydsdgt72")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "qmuvme7f",
    "name": "parent_id",
    "type": "relation",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "collectionId": "r76mw51ydsdgt72",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": null
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("r76mw51ydsdgt72")

  // remove
  collection.schema.removeField("qmuvme7f")

  return dao.saveCollection(collection)
})
