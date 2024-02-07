const mongoose = require('mongoose');

const connectionString = 'mongodb://mongo:27017/imageDB';

function connectToMongoDB() {

    mongoose.connect(connectionString, { useNewUrlParser: true, useUnifiedTopology: true })
        .then(() => console.log('Connected to MongoDB'))
        .catch((error) => console.error('Connection error', error));

}

module.exports = connectToMongoDB;