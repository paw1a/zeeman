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
        error: "Введите правильный размер пленки",
        init: "3"
    },
    {
        id: "waveLength",
        name: "Длина волны",
        error: "Введите правильную длину волны",
        init: "632"
    },
    {
        id: "focalDistance",
        name: "Фокусное расстояние",
        error: "Введите правильное фокусное расстояние",
        init: "100"
    },
    {
        id: "glassesDistance",
        name: "Расстояние между линзами",
        error: "Введите правильное расстояние между линзами",
        init: "10"
    },
    {
        id: "refractionFactor",
        name: "Коэффициент преломления",
        error: "Введите правильный коэффициент преломления",
        init: "1.375"
    },
    {
        id: "magneticInduction",
        name: "Магнитная индукция",
        error: "Введите правильную магнитную индукцию",
        init: "4"
    },
];