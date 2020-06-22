-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 17 Jun 2020 pada 13.36
-- Versi server: 10.4.11-MariaDB
-- Versi PHP: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `fiprosys`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `category`
--

CREATE TABLE `category` (
  `id` int(255) NOT NULL,
  `value` varchar(255) NOT NULL,
  `group_id` int(255) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `category`
--

INSERT INTO `category` (`id`, `value`, `group_id`, `created_at`) VALUES
(1, 'Kemko', 1, '2020-04-18 19:26:13'),
(2, 'Koku', 1, '2020-04-26 09:16:21'),
(3, 'Panjang', 1, '2020-04-26 09:16:21');

-- --------------------------------------------------------

--
-- Struktur dari tabel `category_group`
--

CREATE TABLE `category_group` (
  `id` int(200) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `category_group`
--

INSERT INTO `category_group` (`id`, `name`, `created_at`) VALUES
(1, 'model', '2020-04-18 19:25:15'),
(2, 'lengan', '2020-04-26 09:15:55');

-- --------------------------------------------------------

--
-- Struktur dari tabel `customer`
--

CREATE TABLE `customer` (
  `id` int(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `customer`
--

INSERT INTO `customer` (`id`, `name`, `created_at`) VALUES
(1, 'Muhammad Afif', '2020-05-22 09:26:10'),
(2, 'Aulia Shodiq', '2020-05-22 09:26:10');

-- --------------------------------------------------------

--
-- Struktur dari tabel `producer`
--

CREATE TABLE `producer` (
  `id` int(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `producer`
--

INSERT INTO `producer` (`id`, `name`, `created_at`) VALUES
(1, 'Pak Afif', '2020-04-27 09:32:54');

-- --------------------------------------------------------

--
-- Struktur dari tabel `product`
--

CREATE TABLE `product` (
  `id` int(200) NOT NULL,
  `name` varchar(200) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `product`
--

INSERT INTO `product` (`id`, `name`, `created_at`) VALUES
(1, '01A055', '2020-05-03 07:51:24'),
(2, '03B085', '2020-05-03 07:52:06'),
(3, '03U007', '2020-05-03 07:56:47'),
(5, '06A050', '0000-00-00 00:00:00'),
(6, '06C080', '2020-05-11 06:42:53'),
(7, '24D080', '2020-05-11 06:44:07'),
(10, '07AD053', '0000-00-00 00:00:00'),
(18, '07AD053', '2020-05-14 00:45:03'),
(19, '07AD053', '2020-05-14 00:53:29'),
(20, '05A027', '2020-05-14 00:54:14'),
(21, 'Hangtag Atasan S', '2020-06-01 05:05:48'),
(22, 'Hangtag Atasan M', '2020-06-01 05:05:49'),
(23, 'Hangtag Atasan L', '2020-06-01 05:05:49'),
(24, 'Hangtag Atasan XL', '2020-06-01 05:05:49'),
(25, 'Hangtag Atasan Polos', '2020-06-01 05:06:38'),
(26, 'Hangtag Sirwal S', '2020-06-01 07:25:08'),
(27, 'Hangtag Sirwal M', '2020-06-01 07:25:09'),
(28, 'Hangtag Sirwal L', '2020-06-01 07:25:10'),
(29, 'Hangtag Sirwal XL', '2020-06-01 07:25:11'),
(30, 'Hangtag Sirwal Polos', '2020-06-01 07:26:07'),
(31, 'Care Label 2019', '2020-06-01 07:27:15'),
(32, 'Size Stiker Transparan S', '2020-06-01 07:32:28'),
(33, 'Size Stiker Transparan M', '2020-06-01 07:32:29'),
(34, 'Size Stiker Transparan L', '2020-06-01 07:32:29'),
(35, 'Size Stiker Transparan XL', '2020-06-01 07:32:29'),
(36, 'Size Stiker Transparan 2L', '2020-06-01 07:32:29'),
(37, 'Size Stiker Transparan 3L', '2020-06-01 07:32:30'),
(38, 'Size Stiker Transparan 4L', '2020-06-01 07:32:30'),
(39, 'Label Perisai Woven Atasan', '2020-06-01 07:36:19'),
(40, 'Tali Hangtag Negrito', '2020-06-02 02:05:07'),
(41, 'Label Yoke', '2020-06-02 02:05:07'),
(42, 'Hangtag Peci', '2020-06-02 02:05:07'),
(43, 'Perisai Kulit Glossy', '2020-06-02 02:26:39'),
(44, 'Perisai Kulit Dope', '2020-06-02 02:26:39'),
(45, 'Care Label 2020', '2020-06-02 02:27:59'),
(46, 'Mata Itik Male', '2020-06-02 02:38:31'),
(47, 'Mata Itik Female', '2020-06-02 02:38:33'),
(48, 'Hangtag Peci', '2020-06-02 02:38:35'),
(49, 'Label Peci Woven', '2020-06-02 03:55:20'),
(50, 'Label Perisai Kids Woven', '2020-06-02 03:57:24'),
(51, 'Care Label Kids ', '2020-06-02 03:58:30'),
(52, 'Hangtag Kids', '2020-06-02 04:01:18'),
(53, 'Label Kotak Woven', '2020-06-02 04:03:04'),
(54, 'Hangtag Sandal Persegi', '2020-06-02 04:06:55'),
(55, 'Hangtag Sandal Panjang', '2020-06-02 07:03:33'),
(56, 'Hangtag Sandal Oval', '2020-06-02 07:04:02'),
(57, 'Tali Hitam Gepeng', '2020-06-02 07:15:22'),
(58, 'Tali Biru Bintik Putih', '2020-06-02 07:25:52'),
(59, 'Label Canvas Kecil', '2020-06-02 07:42:56'),
(60, 'Label Canvas Besar', '2020-06-02 07:50:56'),
(61, 'Kancing Putih Susu 18', '2020-06-02 08:27:40'),
(62, 'Kancing Abu Solid 18', '2020-06-02 08:27:40'),
(63, '', '2020-06-02 08:27:41'),
(64, 'Label Size Sirwal S', '2020-06-02 08:46:04'),
(65, 'Label Size Sirwal M', '2020-06-02 08:46:04'),
(66, 'Label Size Sirwal L', '2020-06-02 08:46:05'),
(67, 'Label Size Sirwal XL', '2020-06-02 08:46:05'),
(68, 'Label Size Sirwal 2XL', '2020-06-02 08:46:05'),
(69, 'Label Size Sirwal 3XL', '2020-06-02 08:46:05'),
(70, 'Label Size Sirwal 4XL', '2020-06-02 08:46:06'),
(71, 'Kancing Hitam Glossy 18', '2020-06-02 08:53:38'),
(72, '', '2020-06-03 03:38:39'),
(73, '', '2020-06-03 03:39:04'),
(74, 'Hangtag Koku Versi 1 S', '2020-06-03 03:44:25'),
(75, 'Hangtag Koku Versi 1 M', '2020-06-03 03:44:25'),
(76, 'Hangtag Koku Versi 1 L', '2020-06-03 03:44:26'),
(77, 'Hangtag Koku Versi 1 XL', '2020-06-03 03:44:26'),
(78, 'Label Size Kids 6', '2020-06-03 03:48:21'),
(79, 'Label Size Kids 8', '2020-06-03 03:48:21'),
(80, 'Label Size Kids 10', '2020-06-03 03:48:22'),
(81, 'Label Size Kids 12', '2020-06-03 03:48:22'),
(82, '', '2020-06-03 03:48:22'),
(83, '', '2020-06-03 03:48:22'),
(84, 'Kancing Hitam Glossy 24', '2020-06-03 03:58:58'),
(85, 'Label Persegi Panjang Kulit ', '2020-06-03 03:59:59'),
(86, 'Kancing Bening 20', '2020-06-03 04:01:00'),
(87, 'Kancing Bening Susu 20', '2020-06-03 04:02:10'),
(88, 'Kancing Krim 20', '2020-06-03 04:20:41'),
(89, 'Kancing Krim 20', '2020-06-03 06:26:10'),
(90, 'Hangtag Kurta Versi 1 S', '2020-06-03 07:28:20'),
(91, 'Hangtag Kurta Versi 1 M', '2020-06-03 07:28:21'),
(92, 'Hangtag Kurta Versi 1  L', '2020-06-03 07:28:21'),
(93, 'Hangtag Kurta Versi 1 XL', '2020-06-03 07:28:21'),
(94, 'Hangtag Sirwal Versi 2 Office S', '2020-06-03 07:36:09'),
(95, 'Hangtag Sirwal Versi 2 Office M', '2020-06-03 07:36:09'),
(96, 'Hangtag Sirwal Versi 2 Office L', '2020-06-03 07:36:09'),
(97, 'Hangtag Sirwal Versi 2 Office XL', '2020-06-03 07:36:10'),
(98, 'Kancing Hitam Ungu 20 ', '2020-06-03 08:00:40'),
(99, 'Kancing Abu Pinggir Hitam 20', '2020-06-03 08:00:40'),
(100, 'Kancing Putih Susu 20', '2020-06-03 08:00:40'),
(101, 'Kancing Putih Bening 18', '2020-06-03 08:00:40'),
(102, 'Kancing Putih Susu 18', '2020-06-03 08:00:40'),
(103, 'Kancing Coklat Tua 22', '2020-06-03 08:00:41'),
(104, 'Kancing Abu Pinggir Hitam 22', '2020-06-03 08:00:41'),
(105, 'Kancing Hitam Dope 22 ', '2020-06-03 08:00:41'),
(106, 'Kancing Hitam Glossy 32', '2020-06-03 08:00:41'),
(107, 'Kancing Coklat 32', '2020-06-03 08:00:41'),
(108, 'Hangtag Kemko Versi 2 S', '2020-06-03 08:03:30'),
(109, 'Hangtag Kemko Versi 2 M', '2020-06-03 08:03:30'),
(110, 'Hangtag Kemko Versi 2 L', '2020-06-03 08:03:30'),
(111, 'Hangtag Kemko Versi 2 XL', '2020-06-03 08:03:31'),
(112, 'Hangtag Ghamis Versi 1 S', '2020-06-03 08:06:04'),
(113, 'Hangtag Ghamis Versi 1 M', '2020-06-03 08:06:04'),
(114, 'Hangtag Ghamis Versi 1 L', '2020-06-03 08:06:04'),
(115, 'Hangtag Ghamis Versi 1 XL', '2020-06-03 08:06:04'),
(116, 'Label Size Kids 14', '2020-06-03 08:39:37'),
(117, 'Label Size Kids 4', '2020-06-03 08:39:39'),
(118, 'Label Size Kids 16', '2020-06-03 08:39:39'),
(119, 'Label Size Atasan 3L', '2020-06-03 08:48:54'),
(120, 'Label Size Atasan 4L', '2020-06-03 08:48:55'),
(121, 'Label Size Atasan 2L', '2020-06-03 08:48:56'),
(122, 'Hangtag Sirwal Cargo Versi 1 S', '2020-06-03 09:27:56'),
(123, 'Hangtag Sirwal Cargo Versi 1 M', '2020-06-03 09:27:56'),
(124, 'Hangtag Sirwal Cargo Versi 1 L', '2020-06-03 09:27:56'),
(125, 'Hangtag Sirwal Cargo Versi 1 XL', '2020-06-03 09:27:57'),
(126, 'Hangtag Sirwal Office Versi 1 S', '2020-06-03 09:31:05'),
(127, 'Hangtag Sirwal Office Versi 1 M', '2020-06-03 09:31:06'),
(128, 'Hangtag Sirwal Office Versi 1 L', '2020-06-03 09:31:06'),
(129, 'Hangtag Sirwal Office Versi 1 XL', '2020-06-03 09:31:06'),
(130, 'Hangtag Jubba S 2020', '2020-06-03 09:36:43'),
(131, 'Hangtag Jubba M 2020', '2020-06-03 09:36:43'),
(132, 'Hangtag Jubba L 2020', '2020-06-03 09:36:43'),
(133, 'Hangtag Jubba XL 2020', '2020-06-03 09:36:44'),
(134, 'Af', '2020-06-04 07:40:23');

-- --------------------------------------------------------

--
-- Struktur dari tabel `production`
--

CREATE TABLE `production` (
  `id` int(200) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `quantity` int(200) DEFAULT NULL,
  `product_id` int(200) NOT NULL,
  `producer_id` int(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `production`
--

INSERT INTO `production` (`id`, `created_at`, `quantity`, `product_id`, `producer_id`) VALUES
(1, '2020-05-03 07:51:24', 1200, 1, 1),
(2, '2020-05-03 07:52:06', 10000, 2, 1),
(3, '2020-05-03 07:56:47', 2400, 3, 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `production_status`
--

CREATE TABLE `production_status` (
  `id` int(255) NOT NULL,
  `value_id` int(255) NOT NULL,
  `date` datetime DEFAULT current_timestamp(),
  `production_id` int(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `production_status`
--

INSERT INTO `production_status` (`id`, `value_id`, `date`, `production_id`) VALUES
(1, 1, '2020-05-03 07:51:24', 1),
(2, 2, '2020-05-03 07:52:06', 2),
(3, 1, '2020-05-03 07:56:47', 3),
(4, 2, '0000-00-00 00:00:00', 1),
(5, 2, '2020-05-05 22:52:22', 1),
(6, 1, '2020-06-05 22:52:22', 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `production_status_value`
--

CREATE TABLE `production_status_value` (
  `id` int(255) NOT NULL,
  `value` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `production_status_value`
--

INSERT INTO `production_status_value` (`id`, `value`, `created_at`) VALUES
(1, 'Sampel', '2020-04-27 10:03:14'),
(2, 'Jahit', '2020-04-27 10:03:14');

-- --------------------------------------------------------

--
-- Struktur dari tabel `product_category`
--

CREATE TABLE `product_category` (
  `id` int(255) NOT NULL,
  `product_id` int(255) NOT NULL,
  `category_id` int(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `product_category`
--

INSERT INTO `product_category` (`id`, `product_id`, `category_id`) VALUES
(1, 1, 1),
(2, 2, 2),
(3, 3, 2),
(4, 3, 3),
(5, 5, 1),
(6, 5, 2),
(7, 6, 1),
(8, 6, 2);

-- --------------------------------------------------------

--
-- Struktur dari tabel `sku`
--

CREATE TABLE `sku` (
  `id` int(255) NOT NULL,
  `code` varchar(255) NOT NULL,
  `product_id` int(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `sku`
--

INSERT INTO `sku` (`id`, `code`, `product_id`, `created_at`) VALUES
(13, 'HTS01-S', 21, '2020-06-01 05:05:49'),
(14, 'HTS01-M', 22, '2020-06-01 05:05:49'),
(15, 'HTS01-L', 23, '2020-06-01 05:05:49'),
(16, 'HTS01-XL', 24, '2020-06-01 05:05:50'),
(17, 'HTS-POLOS', 25, '2020-06-01 05:06:38'),
(18, 'HTSIR-S', 26, '2020-06-01 07:25:09'),
(19, 'HTSIR-M', 27, '2020-06-01 07:25:09'),
(20, 'HTSIR-L', 28, '2020-06-01 07:25:11'),
(21, 'HTSIR-XL', 29, '2020-06-01 07:25:11'),
(22, 'HTSIR-POLOS', 30, '2020-06-01 07:26:07'),
(23, 'CALAB2019', 31, '2020-06-01 07:27:15'),
(24, 'SZSTRANS-S', 32, '2020-06-01 07:32:28'),
(25, 'SZSTRANS-M', 33, '2020-06-01 07:32:29'),
(26, 'SZSTRANS-L', 34, '2020-06-01 07:32:29'),
(27, 'SZSTRANS-XL', 35, '2020-06-01 07:32:29'),
(28, 'SZSTRANS-2L', 36, '2020-06-01 07:32:29'),
(29, 'SZSTRANS-3L', 37, '2020-06-01 07:32:30'),
(30, 'SZSTRANS-4L', 38, '2020-06-01 07:32:30'),
(31, 'LPWA', 39, '2020-06-01 07:36:20'),
(32, 'THN2020', 40, '2020-06-02 02:05:07'),
(33, 'MLYOKE', 41, '2020-06-02 02:05:07'),
(34, 'HTP-PRF', 42, '2020-06-02 02:05:07'),
(35, 'PEKULGLOSS', 43, '2020-06-02 02:26:39'),
(36, 'PEKULDOPE', 44, '2020-06-02 02:26:39'),
(37, 'CARLAB2020', 45, '2020-06-02 02:27:59'),
(38, 'MIM', 46, '2020-06-02 02:38:32'),
(39, 'MIF', 47, '2020-06-02 02:38:33'),
(40, 'HTP-PRF', 48, '2020-06-02 02:38:35'),
(41, 'LPWKEC', 49, '2020-06-02 03:55:21'),
(42, 'LPKW', 50, '2020-06-02 03:57:25'),
(43, 'CLKWVN', 51, '2020-06-02 03:58:31'),
(44, 'HTKIDS', 52, '2020-06-02 04:01:18'),
(45, 'LKW', 53, '2020-06-02 04:03:04'),
(46, 'HSANPER', 54, '2020-06-02 04:06:55'),
(47, 'HSANPAN', 55, '2020-06-02 07:03:34'),
(48, 'HSANVAL', 56, '2020-06-02 07:04:03'),
(49, 'THG', 57, '2020-06-02 07:15:24'),
(50, 'TBBP', 58, '2020-06-02 07:25:54'),
(51, 'LCK', 59, '2020-06-02 07:42:57'),
(52, 'LCANBES', 60, '2020-06-02 07:50:56'),
(53, 'KPS18', 61, '2020-06-02 08:27:40'),
(54, 'KAS18', 62, '2020-06-02 08:27:40'),
(55, '', 63, '2020-06-02 08:27:41'),
(56, 'LSSIR-S', 64, '2020-06-02 08:46:04'),
(57, 'LSSIR-M', 65, '2020-06-02 08:46:04'),
(58, 'LSSIR-L', 66, '2020-06-02 08:46:05'),
(59, 'LSSIR-XL', 67, '2020-06-02 08:46:05'),
(60, 'LSSIR-2XL', 68, '2020-06-02 08:46:05'),
(61, 'LSSIR-3XL', 69, '2020-06-02 08:46:06'),
(62, 'LSSIR-4XL', 70, '2020-06-02 08:46:06'),
(63, 'KHG18', 71, '2020-06-02 08:53:38'),
(64, '', 72, '2020-06-03 03:38:40'),
(65, '', 73, '2020-06-03 03:39:04'),
(66, 'HKOV1-S', 74, '2020-06-03 03:44:25'),
(67, 'HKOV1-M', 75, '2020-06-03 03:44:26'),
(68, 'HKOV1-L', 76, '2020-06-03 03:44:26'),
(69, 'HKOV1-XL', 77, '2020-06-03 03:44:26'),
(70, 'LSK-6', 78, '2020-06-03 03:48:21'),
(71, 'LSK-8', 79, '2020-06-03 03:48:21'),
(72, 'LSK-10', 80, '2020-06-03 03:48:22'),
(73, 'LSK-12', 81, '2020-06-03 03:48:22'),
(74, '', 82, '2020-06-03 03:48:22'),
(75, '', 83, '2020-06-03 03:48:22'),
(76, 'KCHGLOSS24', 84, '2020-06-03 03:58:58'),
(77, 'LPPK', 85, '2020-06-03 03:59:59'),
(78, 'KB20', 86, '2020-06-03 04:01:00'),
(79, 'KBS20', 87, '2020-06-03 04:02:11'),
(80, 'KK20', 88, '2020-06-03 04:20:41'),
(81, 'KK20', 89, '2020-06-03 06:26:10'),
(82, 'HKURV1-S', 90, '2020-06-03 07:28:20'),
(83, 'HKURV1-M', 91, '2020-06-03 07:28:21'),
(84, 'HKURV1-L', 92, '2020-06-03 07:28:21'),
(85, 'HKURV1-XL', 93, '2020-06-03 07:28:21'),
(86, 'HSIROFFSV2-S', 94, '2020-06-03 07:36:09'),
(87, 'HSIROFFSV2-M', 95, '2020-06-03 07:36:09'),
(88, 'HSIROFFSV2-L', 96, '2020-06-03 07:36:09'),
(89, 'HSIROFFSV2-XL', 97, '2020-06-03 07:36:10'),
(90, 'KHU20', 98, '2020-06-03 08:00:40'),
(91, 'KAPH20', 99, '2020-06-03 08:00:40'),
(92, 'KAPS20', 100, '2020-06-03 08:00:40'),
(93, 'KPB18', 101, '2020-06-03 08:00:40'),
(94, 'KAPSUS18', 102, '2020-06-03 08:00:40'),
(95, 'KACOKUA22', 103, '2020-06-03 08:00:41'),
(96, 'KAPH20', 104, '2020-06-03 08:00:41'),
(97, 'KHD22', 105, '2020-06-03 08:00:41'),
(98, 'KAHGLOSS32', 106, '2020-06-03 08:00:41'),
(99, 'KACOK 32', 107, '2020-06-03 08:00:41'),
(100, 'HAKEMV2-S', 108, '2020-06-03 08:03:30'),
(101, 'HAKEMV2-M', 109, '2020-06-03 08:03:30'),
(102, 'HAKEMV2-L', 110, '2020-06-03 08:03:31'),
(103, 'HAKEMV2-XL', 111, '2020-06-03 08:03:31'),
(104, 'HANGHAV1-S', 112, '2020-06-03 08:06:04'),
(105, 'HANGHAV1-M', 113, '2020-06-03 08:06:04'),
(106, 'HANGHAV1-L', 114, '2020-06-03 08:06:04'),
(107, 'HANGHAV1-XL', 115, '2020-06-03 08:06:04'),
(108, 'LSK-14', 116, '2020-06-03 08:39:37'),
(109, 'LSK-4', 117, '2020-06-03 08:39:39'),
(110, 'LSK-16', 118, '2020-06-03 08:39:39'),
(111, 'MLBSA-3L', 119, '2020-06-03 08:48:55'),
(112, 'MLBSA-4L', 120, '2020-06-03 08:48:55'),
(113, 'MLBSA-2L', 121, '2020-06-03 08:48:56'),
(114, 'HSCAV1-S', 122, '2020-06-03 09:27:56'),
(115, 'HSCAV1-M', 123, '2020-06-03 09:27:56'),
(116, 'HSCAV1-L', 124, '2020-06-03 09:27:57'),
(117, 'HSCAV1-XL', 125, '2020-06-03 09:27:57'),
(118, 'HSOFFV1-S', 126, '2020-06-03 09:31:05'),
(119, 'HSOFFV1-M', 127, '2020-06-03 09:31:06'),
(120, 'HSOFFV1-L', 128, '2020-06-03 09:31:06'),
(121, 'HSOFFV1-XL', 129, '2020-06-03 09:31:06'),
(122, 'HJ2020-S', 130, '2020-06-03 09:36:43'),
(123, 'HJ2020-M', 131, '2020-06-03 09:36:43'),
(124, 'HJ2020-L', 132, '2020-06-03 09:36:43'),
(125, 'HJ2020-XL', 133, '2020-06-03 09:36:44'),
(126, '12', 134, '2020-06-04 07:40:24');

-- --------------------------------------------------------

--
-- Struktur dari tabel `sku_out`
--

CREATE TABLE `sku_out` (
  `id` int(255) NOT NULL,
  `quantity` int(255) NOT NULL,
  `date` datetime NOT NULL DEFAULT current_timestamp(),
  `sku_id` int(255) NOT NULL,
  `sku_out_group_id` int(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `sku_out`
--

INSERT INTO `sku_out` (`id`, `quantity`, `date`, `sku_id`, `sku_out_group_id`) VALUES
(6, 12, '2020-05-22 02:36:26', 10, 0),
(7, 12, '2020-05-22 03:28:52', 10, 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `sku_out_group`
--

CREATE TABLE `sku_out_group` (
  `id` int(255) NOT NULL,
  `date` datetime NOT NULL,
  `customer_id` int(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `sku_out_group`
--

INSERT INTO `sku_out_group` (`id`, `date`, `customer_id`) VALUES
(1, '2020-05-22 03:28:52', 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `sku_stock`
--

CREATE TABLE `sku_stock` (
  `id` int(255) NOT NULL,
  `sku_id` int(255) NOT NULL,
  `quantity` int(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `sku_stock`
--

INSERT INTO `sku_stock` (`id`, `sku_id`, `quantity`) VALUES
(12, 13, 5396),
(13, 14, 2437),
(14, 15, 934),
(15, 16, 2510),
(16, 17, 0),
(17, 18, 1620),
(18, 19, 4095),
(19, 20, 4300),
(20, 21, 1695),
(21, 22, 100),
(22, 23, 32500),
(23, 24, 415),
(24, 25, 5588),
(25, 26, 8188),
(26, 27, 3928),
(27, 28, 2990),
(28, 29, 3500),
(29, 30, 3585),
(30, 31, 18000),
(31, 32, 42800),
(32, 33, 6490),
(33, 34, 7000),
(34, 35, 3150),
(35, 36, 500),
(36, 37, 5100),
(37, 38, 4412),
(38, 39, 4412),
(39, 40, 7000),
(40, 41, 11800),
(41, 42, 3660),
(42, 43, 3200),
(43, 44, 2500),
(44, 45, 6718),
(45, 46, 4060),
(46, 47, 4600),
(47, 48, 4600),
(48, 49, 100),
(49, 50, 2500),
(50, 51, 7600),
(51, 52, 8400),
(52, 53, 62136),
(53, 54, 8608),
(54, 55, 0),
(55, 56, 3920),
(56, 57, 3814),
(57, 58, 3900),
(58, 59, 2518),
(59, 60, 5445),
(60, 61, 8683),
(61, 62, 7240),
(62, 63, 24896),
(63, 64, 0),
(64, 65, 0),
(65, 66, 1860),
(66, 67, 1869),
(67, 68, 1000),
(68, 69, 3900),
(69, 70, 1672),
(70, 71, 1250),
(71, 72, 1324),
(72, 73, 1488),
(73, 74, 0),
(74, 75, 0),
(75, 76, 36288),
(76, 77, 1250),
(77, 78, 40032),
(78, 79, 37008),
(79, 80, 1728),
(80, 81, 1728),
(81, 82, 1000),
(82, 83, 900),
(83, 84, 2000),
(84, 85, 2000),
(85, 86, 1900),
(86, 87, 1900),
(87, 88, 2900),
(88, 89, 2900),
(89, 90, 33768),
(90, 91, 32544),
(91, 92, 17496),
(92, 93, 1584),
(93, 94, 648),
(94, 95, 34848),
(95, 96, 36504),
(96, 97, 15552),
(97, 98, 55152),
(98, 99, 3844),
(99, 100, 2039),
(100, 101, 2035),
(101, 102, 4000),
(102, 103, 4000),
(103, 104, 800),
(104, 105, 1560),
(105, 106, 1562),
(106, 107, 720),
(107, 108, 1880),
(108, 109, 2503),
(109, 110, 2587),
(110, 111, 5200),
(111, 112, 3140),
(112, 113, 2482),
(113, 114, 700),
(114, 115, 2000),
(115, 116, 1200),
(116, 117, 800),
(117, 118, 350),
(118, 119, 600),
(119, 120, 750),
(120, 121, 400),
(121, 122, 383),
(122, 123, 968),
(123, 124, 994),
(124, 125, 402),
(125, 126, 12);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `category`
--
ALTER TABLE `category`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `category_group`
--
ALTER TABLE `category_group`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `customer`
--
ALTER TABLE `customer`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `producer`
--
ALTER TABLE `producer`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `production`
--
ALTER TABLE `production`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `production_status`
--
ALTER TABLE `production_status`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `production_status_value`
--
ALTER TABLE `production_status_value`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `product_category`
--
ALTER TABLE `product_category`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `sku`
--
ALTER TABLE `sku`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `sku_out`
--
ALTER TABLE `sku_out`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `sku_out_group`
--
ALTER TABLE `sku_out_group`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `sku_stock`
--
ALTER TABLE `sku_stock`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `category`
--
ALTER TABLE `category`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `category_group`
--
ALTER TABLE `category_group`
  MODIFY `id` int(200) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `customer`
--
ALTER TABLE `customer`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `producer`
--
ALTER TABLE `producer`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT untuk tabel `product`
--
ALTER TABLE `product`
  MODIFY `id` int(200) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=135;

--
-- AUTO_INCREMENT untuk tabel `production`
--
ALTER TABLE `production`
  MODIFY `id` int(200) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `production_status`
--
ALTER TABLE `production_status`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT untuk tabel `production_status_value`
--
ALTER TABLE `production_status_value`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `product_category`
--
ALTER TABLE `product_category`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT untuk tabel `sku`
--
ALTER TABLE `sku`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=127;

--
-- AUTO_INCREMENT untuk tabel `sku_out`
--
ALTER TABLE `sku_out`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT untuk tabel `sku_out_group`
--
ALTER TABLE `sku_out_group`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT untuk tabel `sku_stock`
--
ALTER TABLE `sku_stock`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=126;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
