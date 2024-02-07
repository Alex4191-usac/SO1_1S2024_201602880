const AppError = require('../utils/AppError');
const Picture = require('../models/picture-model');

class ImageController {
    static async getAllImages(req, res, next) {
        try{
          const images =  await Picture.find();
          res.status(200).json(images);
        }catch(error){
          next(new AppError(500, 'Error getting images'));
        }
    }

    static async getImage(req, res, next) {
      try {
        const userId = req.params.userId;
        const image = await Picture.findById(userId);
        if (!userId) {
            return next(new AppError(400, 'User ID is required'));
        }
        res.status(200).json(image);
      } catch (error) {
        next(new AppError(500, 'Error getting image'));
      }
        
    }

    static async createImage(req, res, next ) {
      try {
        if(!req.body.date || !req.body.path){
          return next(new AppError(400, 'Date and path are required'));
        }
        const imageData = req.body.path;
        const imageBuffer = Buffer.from(imageData, 'base64');
        const image = new Picture({
          date: req.body.date,
          path: imageBuffer
        
        });
        await image.save();
        res.status(201).json(image);
      } catch (error) {
        next(new AppError(500, 'Error creating image'));
      }    
    }
        
    
}

module.exports = ImageController;