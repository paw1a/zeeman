import {createFieldValidator, requiredValidator} from "./validator";

export type Parameter = {
    id: string,
    name: string,
    error: string,
    init: string
};

export const parameters: Parameter[] = [
    {
        id: "resolution",
        name: "Разрешение изображения",
        error: "Введите правильное разрешение",
        init: "500"
    },
    {
        id: "pictureSize",
        name: "Размер пленки",
        error: "Размер пленки",
        init: "3"
    },
    {
        id: "waveLength",
        name: "Длина волны",
        error: "Размер пленки",
        init: "632"
    },
    {
        id: "focalDistance",
        name: "Фокусное расстояние",
        error: "Размер пленки",
        init: "100"
    },
    {
        id: "glassesDistance",
        name: "Расстояние между линзами",
        error: "Размер пленки",
        init: "10"
    },
    {
        id: "pathDifference",
        name: "Разность хода",
        error: "Размер пленки",
        init: "0"
    },
    {
        id: "refractionFactor",
        name: "Коэффициент преломления",
        error: "Размер пленки",
        init: "1.375"
    },
    {
        id: "magneticInduction",
        name: "Магнитная индукция",
        error: "Размер пленки",
        init: "4"
    },
];