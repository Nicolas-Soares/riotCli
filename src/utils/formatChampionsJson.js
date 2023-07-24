const fs = require("fs")

fs.readFile("champions.json", "utf8", (err, data) => {
  const rawFileData = JSON.parse(data)
  const objFileData = Object.values(rawFileData.data)
  const updatedJson = JSON.stringify(objFileData, null, 2);

  // fs.writeFile("champions.json", updatedJson, "utf8", (err) => {
  //   if (err) console.log("ERROR OCURRED: ", err)
  // })
})
