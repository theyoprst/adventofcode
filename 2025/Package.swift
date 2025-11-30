// swift-tools-version: 6.0
import PackageDescription

let package = Package(
    name: "AdventOfCode2025",
    platforms: [.macOS(.v14)],
    products: [
        .executable(name: "day00", targets: ["Day00"]),
        .executable(name: "day01", targets: ["Day01"]),
    ],
    dependencies: [
        .package(url: "https://github.com/jpsim/Yams.git", from: "5.0.0")
    ],
    targets: [
        .executableTarget(
            name: "Day00",
            path: "00",
            sources: ["Solution.swift"]
        ),
        .testTarget(
            name: "Day00Tests",
            dependencies: ["Day00", "Yams"],
            path: "00",
            sources: ["SolutionTests.swift"],
            resources: [
                .copy("tests.yaml"),
                .copy("input.txt"),
                .copy("input_ex1.txt")
            ]
        ),
        .executableTarget(
            name: "Day01",
            path: "01",
            sources: ["Solution.swift"]
        ),
        .testTarget(
            name: "Day01Tests",
            dependencies: ["Day01", "Yams"],
            path: "01",
            sources: ["SolutionTests.swift"],
            resources: [
                .copy("tests.yaml")
            ]
        ),
    ]
)
