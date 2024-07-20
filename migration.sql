-- phpMyAdmin SQL Dump
-- version 5.1.1deb5ubuntu1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jul 20, 2024 at 10:31 AM
-- Server version: 8.0.37-0ubuntu0.22.04.3
-- PHP Version: 8.3.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `fakehebe`
--

-- --------------------------------------------------------

--
-- Table structure for table `teachers`
--

CREATE TABLE `teachers` (
  `Id` int NOT NULL,
  `Name` varchar(255) DEFAULT NULL,
  `Surname` varchar(255) DEFAULT NULL,
  `DisplayName` varchar(255) NOT NULL,
  `Position` int NOT NULL DEFAULT '1',
  `Description` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `teachers`
--

INSERT INTO `teachers` (`Id`, `Name`, `Surname`, `DisplayName`, `Position`, `Description`) VALUES
(1, 'Krzysztof', 'Dąbrowski', 'Krzysztof Dąbrowski', 1, 'Fizyka'),
(2, 'Maciej', 'Szymczak', 'Maciej Szymczak', 1, 'Chemia'),
(3, 'Arkadiusz', 'Krupa', 'Arkadiusz Krupa', 1, 'Geografia'),
(4, 'Robert', 'Sokołowski', 'Robert Sokołowski', 1, 'Język angielski'),
(5, 'Korneliusz', 'Wojciechowski', 'Korneliusz Wojciechowski', 1, 'Język angielski'),
(6, 'Bogumił', 'Górecki', 'Bogumił Górecki', 1, 'Historia'),
(7, 'Anatol', 'Kamiński', 'Anatol Kamiński', 0, 'Wychowawca'),
(8, NULL, NULL, 'Regina Zawadzka', 2, 'Dyrektor'),
(9, 'Anatol', 'Kamiński', 'Anatol Kamiński', 1, 'Informatyka'),
(10, 'Halina', 'Mazur', 'Halina Mazur', 1, 'Wychowanie do życia w rodzinie'),
(11, 'Irena', 'Baranowska', 'Irena Baranowska', 1, 'Cyfrowa matematyka'),
(12, 'Weronika', 'Wojciechowska', 'Weronika Wojciechowska', 3, 'Pedagog'),
(13, 'Wiktoria', 'Bąk', 'Wiktoria Bąk', 3, 'Pedagog'),
(14, 'Agnieszka', 'Mazurek', 'Agnieszka Mazurek', 3, 'Pedagog'),
(15, 'Waldemar', 'Szulc', 'Waldemar Szulc', 1, 'Historia i teraźniejszość'),
(16, 'Weronika', 'Szulc', 'Weronika Szulc', 1, 'Matematyka'),
(17, 'Bianka', 'Wasilewska', 'Bianka Wasilewska', 1, 'Podstawy przedsiębiorczości'),
(18, 'Anatol', 'Kamiński', 'Anatol Kamiński', 1, 'Zajęcia z wychowawcą'),
(19, 'Łukasz', 'Górecki', 'Łukasz Górecki', 1, 'Język Niemiecki'),
(20, 'Maurycy', 'Chmielewski', 'Maurycy Chmielewski', 1, 'Język Niemiecki'),
(21, 'Remigiusz', 'Kowalski', 'Remigiusz Kowalski', 1, 'Wychowanie fizyczne'),
(22, 'Adam', 'Bąk', 'Adam Bąk', 1, 'Wychowanie fizyczne'),
(23, 'Oktawian', 'Ostrowski', 'Oktawian Ostrowski', 1, 'Biologia'),
(24, 'Olaf', 'Rutkowski', 'Olaf Rutkowski', 1, 'Język polski'),
(25, 'Amadeusz', 'Kamiński', 'Amadeusz Kamiński', 1, 'Religia');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `teachers`
--
ALTER TABLE `teachers`
  ADD PRIMARY KEY (`Id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `teachers`
--
ALTER TABLE `teachers`
  MODIFY `Id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
