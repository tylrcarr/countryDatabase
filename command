db.data.aggregate([{$lookup:{from: "geojson", localField: "alpha3Code", foreignField: "id", as: "geojson"}}, {$project: {"_id": 0, "geometry": 1, "properties": 0, "id": 0, "type":0}}])
