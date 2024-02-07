const mongoose = require('mongoose');
const Schema = mongoose.Schema;

const pictureSchema = new Schema({
    date: {
        type: String,
        required: true
    },
    path: {
        type: Buffer,
        required: true
    }
});

const Picture = mongoose.model('Picture', pictureSchema);

module.exports = Picture;