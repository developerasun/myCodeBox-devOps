import express from 'express'

const app = express()
app.listen(3000, () => console.log("listening at port 3000"))

app.get('/', (req, res) => {
    console.log(req.method)
    res.json({message : "hello there"})
})
