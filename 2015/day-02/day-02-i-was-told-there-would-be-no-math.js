const fs = require("fs")

const data = fs.readFileSync("./day-02-data.txt", { encoding: "utf8" }).split("\n")
data.pop()

const parseLine = (line) => {
  const [l, w, h] = line.split("x")
  return {
    l,
    w,
    h,
  }
}

const transformValues = (obj, cb) => {
  Object.keys(obj).forEach((key) => {
    obj[key] = cb(obj[key])
  })
  return obj
}

const calculateSidesAreas = (parsedLine) => {
  return [
    parsedLine.l * parsedLine.w,
    parsedLine.w * parsedLine.h,
    parsedLine.h * parsedLine.l,
  ]
}

const findSmallestSide = (sidesAreasArr) => {
  return sidesAreasArr.reduce((out, curr) => (out < curr ? out : curr))
}

const wrappingPaperPerBox = (line) => {
  let boxDimensions = transformValues(parseLine(line), parseInt)
  let boxSideAreas = calculateSidesAreas(boxDimensions)
  let smallestSide = findSmallestSide(boxSideAreas)
  return boxSideAreas.reduce((out, curr) => out + 2 * curr, 0) + smallestSide
}

console.log(data.reduce((out, curr) => out + wrappingPaperPerBox(curr), 0))

const shortestPerimeter = (parsedLine) => {
  const sizes = Object.values(parsedLine)
  return 2 * (sizes.reduce((out, curr) => out + curr) - Math.max(...sizes))
}

const packageVolume = (parsedLine) => {
  return Object.values(parsedLine).reduce((out, curr) => out * curr)
}

const ribbonPerBox = (line) => {
  const parsedLine = transformValues(parseLine(line), parseInt)
  return shortestPerimeter(parsedLine) + packageVolume(parsedLine)
}

console.log(data.reduce((out, curr) => out + ribbonPerBox(curr), 0))
