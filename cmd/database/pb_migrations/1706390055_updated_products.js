/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("rrll7kxg0s0ic55")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "dlkfkoaq",
    "name": "category_id",
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
  const collection = dao.findCollectionByNameOrId("rrll7kxg0s0ic55")

  // remove
  collection.schema.removeField("dlkfkoaq")

  return dao.saveCollection(collection)
})
