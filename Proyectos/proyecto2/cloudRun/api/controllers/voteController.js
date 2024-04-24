const Picture = require('../models/vote-model');

class VoteController {
    static async getAllImages(req, res, next) {
        try{
          const images =  await Picture.find();
          res.status(200).json(images);
        }catch(error){
          next(new AppError(500, 'Error getting images'));
        }
    }   

    static async getVotes(req, res, next) {
        try{
          const votes =  await Picture.find().sort({createdAt: -1}).limit(20);
          res.status(200).json(votes);
        }catch(error){
          next(new AppError(500, 'Error getting votes'));
        }
    }
}

    

module.exports = VoteController;