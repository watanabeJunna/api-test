import mongoose, { Schema, Document } from "mongoose"

export interface Struct extends Document {
    Name: string,
}

const StructSchema = new Schema<Struct>({
    Name: { type: String, required: true },
})

export default mongoose.model<Struct>('struct', StructSchema)