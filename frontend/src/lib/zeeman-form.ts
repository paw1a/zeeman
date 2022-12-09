import {createFieldValidator, requiredValidator} from "./validator";

export type Parameter = {
    id: string,
    name: string,
    error: string,
};

export const parameters: Parameter[] = [
    {
        id: "resolution",
        name: "Разрешение изображения",
        error: "Введите правильное разрешение",
    },
    {
        id: "pictureSize",
        name: "Размер пленки",
        error: "Размер пленки",
    },
    {
        id: "waveLength",
        name: "Длина волны",
        error: "Размер пленки",
    },
    {
        id: "focalDistance",
        name: "Фокусное расстояние",
        error: "Размер пленки",
    },
    {
        id: "glassesDistance",
        name: "Расстояние между линзами",
        error: "Размер пленки",
    },
    {
        id: "pathDifference",
        name: "Разность хода",
        error: "Размер пленки",
    },
    {
        id: "refractionFactor",
        name: "Коэффициент преломления",
        error: "Размер пленки",
    },
    {
        id: "magneticInduction",
        name: "Магнитная индукция",
        error: "Размер пленки",
    },
];