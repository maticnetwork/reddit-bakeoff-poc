const args = process.argv 
const request = require('request')
const { blockNumber } = require('../../info.json')
let block

if (args[2]) {
  block = args[2]
} else {
  block = blockNumber
}

let watcher = 'https://mumbai-watcher.api.matic.today/api/v1/header/included/'

async function checkInclusion () {
  let url = watcher + block
  await request(url, { json: true }, (err, res, body) => {
      if (body.status == 'failure') {
        console.log ('\nNot included yet\n---')
        console.log (body)
      } else if (body.status == 'success') {
        console.log ('\ncheckpoint included!\n---')
        console.log (body)
      }
      if (err) { return console.log(err); }
  })
}

checkInclusion()
