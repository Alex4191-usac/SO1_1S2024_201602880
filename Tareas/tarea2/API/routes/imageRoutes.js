const express = require('express')
const router = express.Router()
const imageController = require('../controllers/imageController')

router.get('/', imageController.getAllImages);
router.get('/:id', imageController.getImage);
router.post('/', imageController.createImage);

module.exports = router