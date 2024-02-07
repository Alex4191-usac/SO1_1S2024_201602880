import { useState } from 'react';
import ImageViewer from './components/ImageViewer';
import WebcamCapture from './components/WebcamCapture';
import { ImagesPanel } from './components/Images_panel/ImagesPanel';
import axios from 'axios';


function App() {

  const [capturedImage, setCapturedImage] = useState(null);
  const [images, setImages] = useState([]); // [ {id: 1, image: 'base64', date: '2021-09-01 12:00:00'}, ...
  const [errorsUpload, setErrorsUpload] = useState('');

  const handleCapture = (imageSrc) => {
    setCapturedImage(imageSrc);
  };

  const handleRetake = () => {
    setCapturedImage(null);
  };

  const handleSubmit = (ImageData) => {
    console.log('Submit Image', ImageData);
    ImageData = ImageData.replace('data:image/jpeg;base64,', '');
    const currentDate = new Date().toLocaleString();
    console.log('Image Data:', currentDate, ImageData);
    axios.post('http://localhost:3000/api/images', 
    { date: currentDate, path: ImageData })

    .then((response) => {
      console.log('Response:', response);
      setErrorsUpload('Image uploaded successfully');
      handleListImages();
    }).
    catch((error) => {
      console.error('Error:', error);
      setErrorsUpload('Error uploading image');
    });
    
   

    setTimeout(() => {
      setErrorsUpload('');
    }, 3000);

  };

  const handleListImages = () => {
    console.log('List Images', images.length);

    axios.get('http://localhost:3000/api/images',
    { headers: { 'Content-Type': 'application/json' }})
    .then((response) => {
      console.log('Response:', response);
      setImages(response.data);
    }).
    catch((error) => {
      console.error('Error:', error);
    });
    
    
  }


  return (
    <section className='min-h-screen flex bg-slate-50 '>
      <div className='p-10 w-2/3  '>
        <h1 className='text-4xl text-center'>Screen WebCam Uploader</h1>
        <p className='text-xl text-gray-700 text-center'> Upload your webcam image to the server</p>
        {capturedImage ? (
          
          <ImageViewer capturedImage={capturedImage} onRetake={handleRetake} onSendData={handleSubmit} />
        ) : (
        
          <WebcamCapture onCapture={handleCapture} />
        )}
       {errorsUpload!=='' &&
          <div 
            className=' bg-green-700 text-center rounded-lg p-3 mt-5 text-white font-bold text-lg w-1/2'>
            <p>{errorsUpload}</p>
          </div>
       }
      </div>
      <div className='w-1/3 p-10 border rounded-lg border-gray-500 '>
        <h1 className='text-4xl text-center'>My Photos</h1>
        <p className='text-xl text-gray-700 text-center'>currrent images uploaded on your folder</p>
        <button
          className=' mt-4 px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600'
          onClick={handleListImages}
        >
          Refresh
          {/*Call the list of Images uploaded */} 
          
        </button>
        <ImagesPanel album={images} />
      </div>
     
    </section>
  )
}

export default App

/* */