import { MongoError } from "mongodb"
import mongoose from "mongoose"
import User, { IUser } from "./User"

const options = {
    useNewUrlParser: true
}

const callback = (err: MongoError) => {
    if (err) {
        throw err
    }
}

mongoose.connect('mongodb://localhost:27017/test', options, callback)

User.find(<IUser | {}>{}, (err: Error, res: any) => {
    if (err) {
        throw err
    }

    console.log(res)

    mongoose.disconnect()
})