

const ImageViewer = ({ capturedImage, onRetake, onSendData }) => {
  return (
    <div className="mb-4 mt-4">
      <div className="flex flex-col justify-center items-center">
        <img src={capturedImage} alt="Captured"  />
        <div className="flex w-1/2">
        <button
          className=" w-1/2 m-2 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
          onClick={onRetake}
        >
          Retake Photo
        </button>
        <button
          className=" w-1/2 m-2 px-4 py-2 bg-green-800 text-white rounded hover:bg-blue-600"
          onClick={() => onSendData(capturedImage)}
        >
          Send to the Server
        </button>
        </div>

      </div>
    </div>
  );
};

export default ImageViewer;
