import { useState } from 'react'
import { Buffer } from 'buffer'

export const ImagesPanel = ({album}) => {
  const [GetImage, setImage] = useState('')

  const handleViewImage = (image) => {
    const base64Data = Buffer.from(image).toString('base64');
    const imageUrl = `data:image/jpeg;base64,${base64Data}`
    setImage(imageUrl)

  }

  return (
    <div className='border '>
        <p className=' p-2 font-semibold text-gray-700 text-lg'>Images uploaded: </p>
        <div className=' p-3 overflow-y-auto h-80 '>
          {
            album.length > 0 && album.map((image, index) => (
              <div key={index} className=' flex items-center p-3 '>
               <button
                  onClick={() => handleViewImage(image.path)}
                  className='bg-blue-500 text-white px-4 py-2 rounded-lg'>View</button>
                  
                <p className='m-2 text-xl text-gray-700'>image uploaded at: {image.date}</p>
              </div>
            ))
          }

        </div>
        <p className=' p-2 font-semibold text-gray-700 text-lg'>Image Viewer: </p>
        <div className='border-4 rounded-xl p-3 w-full h-60 border-grey-500'>
            
            {
              
              GetImage && (
                <img src={GetImage} alt='image of view' className='max-w-full max-h-full' />
              )
              
            }
        </div>
    </div>


  )
}
