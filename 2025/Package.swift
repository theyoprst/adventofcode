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
        .executable(name: "aoc", targets: ["aoc"]),
    ],
    dependencies: [
        .package(url: "https://github.com/jpsim/Yams.git", from: "5.0.0"),
        .package(url: "https://github.com/apple/swift-argument-parser.git", exact: "1.6.2"),
        .package(url: "https://github.com/apple/swift-collections.git", .upToNextMajor(from: "1.3.0")),
    ],
    targets: [
        // Shared utilities
        .target(
            name: "AOCUtilities",
            path: "AOCUtilities"
        ),

        // Single executable with all days
        .executableTarget(
            name: "aoc",
            dependencies: [
                "AOCUtilities",
                .product(name: "ArgumentParser", package: "swift-argument-parser"),
                .product(name: "Collections", package: "swift-collections"),
            ],
            path: "Sources/aoc",
            resources: [
                .copy("../../Resources/01"),
                .copy("../../Resources/02"),
                .copy("../../Resources/03"),
                .copy("../../Resources/04"),
                .copy("../../Resources/05"),
                .copy("../../Resources/06"),
                .copy("../../Resources/07"),
                .copy("../../Resources/08"),
                .copy("../../Resources/09"),
            ]
        ),

        // Single test target
        .testTarget(
            name: "AOCTests",
            dependencies: [
                "aoc",
                "AOCUtilities",
                "Yams",
                .product(name: "Collections", package: "swift-collections"),
            ],
            path: "Tests/AOCTests",
            resources: [
                .copy("../../Resources/01"),
                .copy("../../Resources/02"),
                .copy("../../Resources/03"),
                .copy("../../Resources/04"),
                .copy("../../Resources/05"),
                .copy("../../Resources/06"),
                .copy("../../Resources/07"),
                .copy("../../Resources/08"),
                .copy("../../Resources/09"),
            ]
        ),
    ]
)
