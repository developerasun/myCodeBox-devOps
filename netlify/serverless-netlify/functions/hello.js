exports.handler = async (event, context) => {
    return { 
        statusCode : 200, 
        body : JSON.stringify({ message : "hello serverless, I'm coming from backend" })
    }
}