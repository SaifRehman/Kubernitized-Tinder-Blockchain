require('dotenv').config({ path: ".env-node1" });
var geodist = require('geodist')
let lotion = require('lotion')
let app = lotion({
  genesis: './genesis.json',
  tendermintPort: 30090,
  initialState: { messages: [] },
  p2pPort: 30091,
  logTendermint: true,
  keys: 'privkey0.json',
  createEmptyBlocks: true,
  peers: ['192.168.1.102:30093']
})
app.use((state, tx, chainInfo) => {
  if (tx[0].gender != tx[1].gender && Math.abs(tx[0].age - tx[1].age) <= 4) {
    var dist = geodist({ lat: tx[0].userlocation.Lat, lon: tx[0].userlocation.Long }, { lat: tx[1].userlocation.Lat, lon: tx[1].userlocation.Long }, { exact: true, unit: 'km' })
    if (dist < 1.0) {
      state.messages.push({
        userone: {
          name: tx[0].name,
          gender: tx[0].gender,
          age: tx[0].age
        },
        usertwo: {
          name: tx[1].name,
          gender: tx[1].gender,
          age: tx[1].age
        }
      })
    }
  }
})
app.listen(3000).then(({ GCI }) => {
  console.log(GCI)
})