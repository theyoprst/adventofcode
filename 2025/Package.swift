// swift-tools-version: 6.0
import PackageDescription

let package = Package(
    name: "AdventOfCode2025",
    platforms: [
        .macOS(.v13),
        .iOS(.v16),
        .tvOS(.v16),
        .watchOS(.v9),
    ],
    products: [
        .executable(name: "day00", targets: ["Day00"]),
        .executable(name: "day01", targets: ["Day01"]),
        .executable(name: "day02", targets: ["Day02"]),
        .executable(name: "day03", targets: ["Day03"]),
        .executable(name: "day04", targets: ["Day04"]),
    ],
    dependencies: [
        .package(url: "https://github.com/jpsim/Yams.git", from: "5.0.0"),
    ],
    targets: [
        .target(
            name: "AOCUtilities",
            path: "AOCUtilities",
        ),
        .target(
            name: "AOCTestSupport",
            dependencies: ["Yams", "AOCUtilities"],
            path: "AOCTestSupport",
        ),
        .executableTarget(
            name: "Day00",
            dependencies: ["AOCUtilities"],
            path: "00",
            exclude: ["SolutionTests.swift", "tests.yaml", "input.txt", "input_ex1.txt"],
            sources: ["Solution.swift"],
        ),
        .testTarget(
            name: "Day00Tests",
            dependencies: ["Day00", "AOCTestSupport"],
            path: "00",
            exclude: ["Solution.swift"],
            sources: ["SolutionTests.swift"],
            resources: [
                .copy("tests.yaml"),
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
        .executableTarget(
            name: "Day01",
            dependencies: ["AOCUtilities"],
            path: "01",
            exclude: ["SolutionTests.swift", "tests.yaml", "input.txt", "input_ex1.txt", "part1.html", "part1.md"],
            sources: ["Solution.swift"],
        ),
        .executableTarget(
            name: "Day02",
            dependencies: ["AOCUtilities"],
            path: "02",
            exclude: ["SolutionTests.swift", "tests.yaml"],
            sources: ["Solution.swift"],
        ),
        .executableTarget(
            name: "Day03",
            dependencies: ["AOCUtilities"],
            path: "03",
            exclude: ["SolutionTests.swift", "tests.yaml"],
            sources: ["Solution.swift"],
        ),
        .executableTarget(
            name: "Day04",
            dependencies: ["AOCUtilities"],
            path: "04",
            exclude: ["SolutionTests.swift", "tests.yaml"],
            sources: ["Solution.swift"],
        ),
        .testTarget(
            name: "Day01Tests",
            dependencies: ["Day01", "AOCTestSupport"],
            path: "01",
            exclude: ["Solution.swift", "part1.html", "part1.md"],
            sources: ["SolutionTests.swift"],
            resources: [
                .copy("tests.yaml"),
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
        .testTarget(
            name: "Day02Tests",
            dependencies: ["Day02", "AOCTestSupport"],
            path: "02",
            exclude: ["Solution.swift"],
            sources: ["SolutionTests.swift"],
            resources: [
                .copy("tests.yaml"),
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
        .testTarget(
            name: "Day03Tests",
            dependencies: ["Day03", "AOCTestSupport"],
            path: "03",
            exclude: ["Solution.swift"],
            sources: ["SolutionTests.swift"],
            resources: [
                .copy("tests.yaml"),
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
        .testTarget(
            name: "Day04Tests",
            dependencies: ["Day04", "AOCTestSupport"],
            path: "04",
            exclude: ["Solution.swift"],
            sources: ["SolutionTests.swift"],
            resources: [
                .copy("tests.yaml"),
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
    ]
)
