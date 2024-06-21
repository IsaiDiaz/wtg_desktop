import ToTry from "../../components/ToTryComponent"

function Hello() {
    const times = 15;
    return (
        <div>
            <h1>Hello</h1>
            {[...Array(times)].map((_, index) => (
                <ToTry key={index} />
            ))}
        </div>
    )
}

export default Hello