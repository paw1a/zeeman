import axios, {AxiosRequestConfig} from "axios";

type ZeemanAPIRequest = {
    resolution: number, // pix
    pictureSize: number, // mm
    waveLength: number, // nm
    focalDistance: number, // mm
    glassesDistance: number,
    pathDifference: number,
    refractionFactor: number,
    magneticInduction: number
};

export default function zeemanAPIRequest(dto: ZeemanAPIRequest) {
    console.log("abc");
    const config = {
        headers: {
            'Access-Control-Allow-Origin': '*'
        },
    };
    fetch(`http://localhost:8080/api/interferometer`, {
        method: "post",
        body: JSON.stringify(dto)
    }).then(resp => resp.blob())
        .then(resp => {
            console.log(resp);
            let img = URL.createObjectURL(resp);
            document.getElementById("aaa").setAttribute("src", img);
        }).catch(err => console.error(err));
};

/*
{
    "resolution": 500,
    "pictureSize": 3,
    "waveLength": 632,
    "focalDistance": 100,
    "glassesDistance": 10,
    "pathDifference": 0,
    "refractionFactor": 1.375,
    "reflectionFactor": 0.7,
    "magneticInduction": 3
}
 */