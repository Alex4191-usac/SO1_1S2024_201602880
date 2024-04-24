const mongoose = require('mongoose');

const connectionString = 'mongodb://34.28.96.178:27017/sopes1p2';

function connectToMongoDB() {

    mongoose.connect(connectionString)
        .then(() => console.log('Connected to MongoDB'))
        .catch((error) => console.error('Connection error', error));

}

module.exports = connectToMongoDB;