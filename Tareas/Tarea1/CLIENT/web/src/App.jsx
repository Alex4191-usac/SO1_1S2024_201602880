import { Suspense, useState, lazy } from "react"


const User = lazy(() => import("./User.jsx"));

export default function App() {
  const [getDataUser, setDataUser] = useState([])
  const [showData, setShowData] = useState(false);

  const getData = async () => {
    const response = await fetch("http://localhost:8080/data")
    const data = await response.json()
    setDataUser(data.message)
  }

  const showDataUser = () => {
    setShowData(!showData)
    getData()
  }



  return (
    <section className="home-screen">
      <h1>Load Data  from Go Server</h1>
      <div className="form-home-screen">
        <button 
          className="btn-server"
          onClick={showDataUser}
        >
          {showData ? "Hide Creedentials" : "Show Creedentials"}
        </button>
        <p>Current Profile:</p>
        <div className="profile">
          {showData && (
            <Suspense fallback={<Loading />}>
              <User data={getDataUser} />
            </Suspense>
          ) }
        </div>
      </div>
    </section>
  )
}

function Loading() {
  return <p>Loading...</p>;
}






