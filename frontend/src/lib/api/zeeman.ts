export enum MagneticFieldDir {
    Parallel,
    Perpendicular
};


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