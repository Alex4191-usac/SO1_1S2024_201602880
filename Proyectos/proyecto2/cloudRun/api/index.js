const express = require('express')
const AppError = require('./utils/AppError');
const cors = require('cors')
const bodyParser = require('body-parser')
const db = require('./utils/db/index')

const app = express()
const port = process.env.PORT || 3000; 

app.use(cors())
app.use(bodyParser.json({ limit: '10mb' }));
app.use(bodyParser.urlencoded({ limit: '10mb', extended: true }));

db();


const imageRoutes = require('./routes/voteRoutes');

app.use('/api/votes', imageRoutes);

app.use((req, res, next)=> {
    const err = new AppError(404, `Can't find ${req.originalUrl} on this server!`);
    next(err);
});

app.use((err, req, res, next) => {
    err.statusCode = err.statusCode || 500;
    err.status = err.status || 'error';
  
    res.status(err.statusCode).json({
      status: err.status,
      message: err.message,
    });
});


app.listen(port, () => {
    console.log(`Server is running on port ${port}`)
});