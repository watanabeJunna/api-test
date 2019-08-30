import { MongoError } from "mongodb"
import mongoose from "mongoose"
import Account, { Struct } from "./Struct"

mongoose.connect('mongodb://localhost:27017/test', {
    useNewUrlParser: true
}, (err: MongoError) => {
    if (err) throw err
})

Account.find(<Struct | {}>{}, (err: Error, res: any) => {
    if (err) throw err
    console.log(res)
    mongoose.disconnect()
})