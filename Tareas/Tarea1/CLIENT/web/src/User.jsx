const User = ({ data }) => {
    return (
        <div>
            <p>Name: {data.name}</p>
            <p>Id: {data.id}</p>
        </div>
    )
}

export default User