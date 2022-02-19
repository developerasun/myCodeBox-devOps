const express = require('express')
const app = express()

app.get('/', (req, res) => {
    res.json([
        {
        'name' : 'not jake', 
        }, 
        {
        'name' : 'brian', 
        }, 
        {
        'name' : 'paul', 
        }, 
    ])
})

app.listen(4000, ()=>console.log('listen at 4000'))