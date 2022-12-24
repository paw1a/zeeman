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
        name: "Разрешение изображения, пкс",
        error: "Введите правильное разрешение",
        init: "500"
    },
    {
        id: "pictureSize",
        name: "Размер пленки, мм",
        error: "Введите правильный размер пленки",
        init: "3"
    },
    {
        id: "waveLength",
        name: "Длина волны, нм",
        error: "Введите правильную длину волны",
        init: "632"
    },
    {
        id: "focalDistance",
        name: "Фокусное расстояние, мм",
        error: "Введите правильное фокусное расстояние",
        init: "100"
    },
    {
        id: "glassesDistance",
        name: "Расстояние между линзами, мм",
        error: "Введите правильное расстояние между линзами",
        init: "10"
    },
    {
        id: "refractionFactor",
        name: "Коэффициент преломления",
        error: "Введите правильный коэффициент преломления",
        init: "0.5"
    },
    {
        id: "magneticInduction",
        name: "Магнитная индукция, мТл",
        error: "Введите правильную магнитную индукцию",
        init: "4"
    },
];