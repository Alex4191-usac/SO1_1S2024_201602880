const mongoose = require('mongoose');
const Schema = mongoose.Schema;

const voteSchema = new Schema({
    name: {
        type: String
    },
    album: {
        type: String
    },
    year: {
        type: String
    },
    rank : {
        type: String
    }
});

const Votos = mongoose.model('Votos', voteSchema);

module.exports = Votos;