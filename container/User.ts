import mongoose, { Schema, Document } from "mongoose"

export interface IUser extends Document {
    Name: string,
}

const UserSchema = new Schema<IUser>({
    Name: { type: String, required: true },
}, { collection: 'user' })

export default mongoose.model<IUser>('user', UserSchema)