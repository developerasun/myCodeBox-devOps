const functions = require("firebase-functions");
const express = require("express");
const app = express();

app.get("/", (req, res)=> {
  res.send("<h1>Hello Firebase</h1>");
});

app.listen(3000, ()=>console.log("listening at 3000"));

exports.app = functions.https.onRequest(app);
// // Create and Deploy Your First Cloud Functions
// // https://firebase.google.com/docs/functions/write-firebase-functions
//
// exports.helloWorld = functions.https.onRequest((request, response) => {
//   functions.logger.info("Hello logs!", {structuredData: true});
//   response.send("Hello from Firebase!");
// });
