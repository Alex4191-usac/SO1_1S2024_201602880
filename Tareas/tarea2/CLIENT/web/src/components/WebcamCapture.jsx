import { useRef, useState } from 'react';
import Webcam from 'react-webcam';

const WebcamCapture = ({ onCapture }) => {
  const webcamRef = useRef(null);

  const captureImage = () => {
    const imageSrc = webcamRef.current.getScreenshot();
    onCapture(imageSrc);
  };

  return (
    <div className='pt-5 '>
      <div className='flex flex-col justify-center items-center'>
        <Webcam
          audio={false}
          ref={webcamRef}
          height={300}
          screenshotFormat="image/jpeg"
          className="rounded mb-4"
        />
        <button
          className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600"
          onClick={captureImage}
        >
          Capture Photo
        </button>
      </div>
    </div>
  );
};

export default WebcamCapture;