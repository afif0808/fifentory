-- phpMyAdmin SQL Dump
-- version 4.9.0.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 23 Jun 2020 pada 06.38
-- Versi server: 10.4.6-MariaDB
-- Versi PHP: 7.3.8

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
(1, 'PT Ambassador Garmindo', '2020-06-17 00:00:00'),
(2, 'PT Efrata Retailindo', '2020-06-17 20:46:07'),
(3, 'Pak Dani', '2020-06-17 20:46:38'),
(4, 'Pak Egi', '2020-06-17 20:46:38'),
(5, 'Felix', '2020-06-17 20:46:56'),
(6, 'Pak Haidar', '2020-06-17 20:46:56'),
(7, 'Bu Heri', '2020-06-17 20:47:08'),
(8, 'Pak Deden', '2020-06-17 20:47:08'),
(9, 'Pak Gobi', '2020-06-18 19:46:23'),
(10, 'Pak Ami', '2020-06-18 19:46:23'),
(11, 'Pak Haji Abang', '2020-06-18 19:55:59'),
(12, 'Samase House', '2020-06-23 10:12:09'),
(13, 'Samase Pasbar', '2020-06-23 10:13:23');

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
(90, 'Hangtag Kurta Versi 2 S', '2020-06-03 07:28:20'),
(91, 'Hangtag Kurta Versi 2 M', '2020-06-03 07:28:21'),
(92, 'Hangtag Kurta Versi 2  L', '2020-06-03 07:28:21'),
(93, 'Hangtag Kurta Versi 2 XL', '2020-06-03 07:28:21'),
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
(130, 'Hangtag Jubba S 2020', '2020-06-03 09:36:43'),
(131, 'Hangtag Jubba M 2020', '2020-06-03 09:36:43'),
(132, 'Hangtag Jubba L 2020', '2020-06-03 09:36:43'),
(133, 'Hangtag Jubba XL 2020', '2020-06-03 09:36:44'),
(134, 'Af', '2020-06-04 07:40:23'),
(135, 'Label Size Atasan S', '2020-06-17 11:45:18'),
(136, 'Label Size Atasan M', '2020-06-17 11:45:18'),
(137, 'Label Size Atasan L', '2020-06-17 11:45:18'),
(138, 'Label Size Atasan XL', '2020-06-17 11:45:18'),
(139, 'Kancing Coklat Pinggir Hitam 18', '2020-06-17 11:46:10'),
(140, 'Label Segi Enam Woven', '2020-06-17 11:47:14'),
(141, 'Hangtag Sirwal Cargo 2 S', '2020-06-17 11:50:21'),
(142, 'Hangtag Sirwal Cargo 2 M', '2020-06-17 11:50:21'),
(143, 'Hangtag Sirwal Cargo 2 L', '2020-06-17 11:50:21'),
(144, 'Hangtag Sirwal Cargo 2 XL', '2020-06-17 11:50:21'),
(145, 'Hangtag Sirwal Cargo 2 POLOS', '2020-06-17 11:50:21'),
(146, 'Hangtag Sirwal Jogger S', '2020-06-17 11:53:00'),
(147, 'Hangtag Sirwal Jogger M', '2020-06-17 11:53:00'),
(148, 'Hangtag Sirwal Jogger L', '2020-06-17 11:53:00'),
(149, 'Hangtag Sirwal Jogger XL', '2020-06-17 11:53:00'),
(150, 'Hangtag Sirwal Jogger Polos', '2020-06-17 11:53:00'),
(151, 'Polibag Sirwal 35x45 2020', '2020-06-17 11:56:55'),
(152, 'Hangtag Koku Versi 2 S', '2020-06-17 12:00:20'),
(153, 'Hangtag Koku Versi 2 M', '2020-06-17 12:00:20'),
(154, 'Hangtag Koku Versi 2 L', '2020-06-17 12:00:20'),
(155, 'Hangtag Koku Versi 2 XL', '2020-06-17 12:00:20'),
(156, 'Hangtag Kurta Versi 1 S', '2020-06-17 12:05:15'),
(157, 'Hangtag Kurta Versi 1 M', '2020-06-17 12:05:15'),
(158, 'Hangtag Kurta Versi 1 L', '2020-06-17 12:05:15'),
(159, 'Hangtag Kurta Versi 1 XL', '2020-06-17 12:05:15'),
(160, 'Hangtag Kemko Versi 1 Besar S', '2020-06-17 13:08:27'),
(161, 'Hangtag Kemko Versi 1 Besar M', '2020-06-17 13:08:27'),
(162, 'Hangtag Kemko Versi 1 Besar L', '2020-06-17 13:08:27'),
(163, 'Hangtag Kemko Versi 1 Besar XL', '2020-06-17 13:08:27'),
(164, 'Hangtag Kemko Versi 1 Besar 2L', '2020-06-17 13:08:27'),
(165, 'Hangtag Kemko Versi 1 Besar 3L', '2020-06-17 13:08:28'),
(166, 'Hangtag Kemko Versi 1 Besar 4L', '2020-06-17 13:08:28'),
(167, 'Hangtag Kemko Office S', '2020-06-17 13:14:05'),
(168, 'Hangtag Kemko Office M', '2020-06-17 13:14:05'),
(169, 'Hangtag Kemko Office L', '2020-06-17 13:14:05'),
(170, 'Hangtag Kemko Office XL', '2020-06-17 13:14:06'),
(171, 'Hangtag Kemko Office 2L', '2020-06-17 13:14:06'),
(172, 'Hangtag Sirwal Office Versi 1 S', '2020-06-17 13:23:27'),
(173, 'Hangtag Sirwal Office Versi 1 M', '2020-06-17 13:23:27'),
(174, 'Hangtag Sirwal Office Versi 1 L', '2020-06-17 13:23:27'),
(175, 'Hangtag Sirwal Office Versi 1 XL', '2020-06-17 13:23:27'),
(176, '', '2020-06-17 13:23:28'),
(177, 'Kancing Hitam Corak Putih 18', '2020-06-17 13:29:25'),
(178, 'Kancing Navy 18', '2020-06-17 13:30:13'),
(179, 'Kancing Abu Pinggir Putih 18', '2020-06-17 13:30:42'),
(180, 'Kancing Coklat Pinggir Kream 18', '2020-06-17 13:31:19'),
(181, 'Kancing Coklat Tua 18', '2020-06-17 13:31:39'),
(182, 'Kancing Krem 18', '2020-06-17 13:36:52'),
(183, 'Kancing Coklat Krem Kotor 18', '2020-06-17 13:38:18'),
(184, 'Plastik Sin 30x40', '2020-06-18 12:35:33'),
(185, 'Stopper', '2020-06-18 12:49:26'),
(186, 'Tali Karet', '2020-06-18 12:49:27');

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
(82, 'HKURV2-S', 90, '2020-06-03 07:28:20'),
(83, 'HKURV2-M', 91, '2020-06-03 07:28:21'),
(84, 'HKURV2-L', 92, '2020-06-03 07:28:21'),
(85, 'HKURV2-XL', 93, '2020-06-03 07:28:21'),
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
(122, 'HJ2020-S', 130, '2020-06-03 09:36:43'),
(123, 'HJ2020-M', 131, '2020-06-03 09:36:43'),
(124, 'HJ2020-L', 132, '2020-06-03 09:36:43'),
(125, 'HJ2020-XL', 133, '2020-06-03 09:36:44'),
(126, '12', 134, '2020-06-04 07:40:24'),
(127, 'LSA-S', 135, '2020-06-17 11:45:18'),
(128, 'LSA-M', 136, '2020-06-17 11:45:18'),
(129, 'LSA-L', 137, '2020-06-17 11:45:18'),
(130, 'LSA-XL', 138, '2020-06-17 11:45:19'),
(131, 'KCPH18', 139, '2020-06-17 11:46:10'),
(132, 'LSEW', 140, '2020-06-17 11:47:14'),
(133, 'HSC2-S', 141, '2020-06-17 11:50:21'),
(134, 'HSC2-M', 142, '2020-06-17 11:50:21'),
(135, 'HSC2-L', 143, '2020-06-17 11:50:21'),
(136, 'HSC2-XL', 144, '2020-06-17 11:50:21'),
(137, 'HSC2-POLOS', 145, '2020-06-17 11:50:21'),
(138, 'HSJOGG-S', 146, '2020-06-17 11:53:00'),
(139, 'HSJOGG-M', 147, '2020-06-17 11:53:00'),
(140, 'HSJOGG-L', 148, '2020-06-17 11:53:00'),
(141, 'HSJOGG-XL', 149, '2020-06-17 11:53:00'),
(142, 'HSJOGG-POLOS', 150, '2020-06-17 11:53:00'),
(143, 'PS35452020', 151, '2020-06-17 11:56:55'),
(144, 'HKOKUV2-S', 152, '2020-06-17 12:00:20'),
(145, 'HKOKUV2-M', 153, '2020-06-17 12:00:20'),
(146, 'HKOKUV2-L', 154, '2020-06-17 12:00:20'),
(147, 'HKOKUV2-XL', 155, '2020-06-17 12:00:20'),
(148, 'HKURV1-S', 156, '2020-06-17 12:05:15'),
(149, 'HKURV1-M', 157, '2020-06-17 12:05:15'),
(150, 'HKURV1-L', 158, '2020-06-17 12:05:15'),
(151, 'HKURV1-XL', 159, '2020-06-17 12:05:15'),
(152, 'HKV1B-S', 160, '2020-06-17 13:08:27'),
(153, 'HKV1B-M', 161, '2020-06-17 13:08:27'),
(154, 'HKV1B-L', 162, '2020-06-17 13:08:27'),
(155, 'HKV1B-XL', 163, '2020-06-17 13:08:27'),
(156, 'HKV1B-2L', 164, '2020-06-17 13:08:27'),
(157, 'HKV1B-3L', 165, '2020-06-17 13:08:28'),
(158, 'HKV1B-4L', 166, '2020-06-17 13:08:28'),
(159, 'HKOFF-S', 167, '2020-06-17 13:14:05'),
(160, 'HKOFF-M', 168, '2020-06-17 13:14:05'),
(161, 'HKOFF-L', 169, '2020-06-17 13:14:06'),
(162, 'HKOFF-XL', 170, '2020-06-17 13:14:06'),
(163, 'HKOFF-2L', 171, '2020-06-17 13:14:06'),
(164, 'HSOF1-S', 172, '2020-06-17 13:23:27'),
(165, 'HSOF1-M', 173, '2020-06-17 13:23:27'),
(166, 'HSOF1-L', 174, '2020-06-17 13:23:27'),
(167, 'HSOF1-XL', 175, '2020-06-17 13:23:27'),
(168, '', 176, '2020-06-17 13:23:28'),
(169, 'KHCOKP18', 177, '2020-06-17 13:29:25'),
(170, 'KHNAVY18', 178, '2020-06-17 13:30:13'),
(171, 'KABPUT18', 179, '2020-06-17 13:30:42'),
(172, 'KCPKREAM18', 180, '2020-06-17 13:31:19'),
(173, 'KACOKUA18', 181, '2020-06-17 13:31:39'),
(174, 'KKOT', 182, '2020-06-17 13:36:53'),
(175, 'KACOKKOT18', 183, '2020-06-17 13:38:18'),
(176, 'PS30X40', 184, '2020-06-18 12:35:34'),
(177, 'STP', 185, '2020-06-18 12:49:26'),
(178, 'TKAR', 186, '2020-06-18 12:49:27');

-- --------------------------------------------------------

--
-- Struktur dari tabel `sku_in`
--

CREATE TABLE `sku_in` (
  `id` int(200) NOT NULL,
  `quantity` int(200) NOT NULL,
  `sku_id` int(200) NOT NULL,
  `date` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Struktur dari tabel `sku_in_group`
--

CREATE TABLE `sku_in_group` (
  `id` int(200) NOT NULL,
  `supplier_id` int(200) NOT NULL,
  `date` date NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

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
(10, 310, '2020-06-09 00:00:00', 13, 4),
(11, 310, '2020-06-09 00:00:00', 14, 4),
(12, 620, '2020-06-09 00:00:00', 15, 4),
(13, 620, '2020-06-09 00:00:00', 16, 4),
(14, 584, '2020-06-09 00:00:00', 24, 4),
(15, 600, '2020-06-09 00:00:00', 25, 4),
(16, 1197, '2020-06-09 00:00:00', 26, 4),
(17, 1200, '2020-06-09 00:00:00', 27, 4),
(18, 3400, '2020-06-09 00:00:00', 31, 4),
(19, 3400, '2020-06-09 00:00:00', 37, 4),
(20, 3456, '2020-06-09 00:00:00', 63, 4),
(21, 600, '2020-06-09 00:00:00', 127, 4),
(22, 600, '2020-06-09 00:00:00', 128, 4),
(23, 850, '2020-06-09 00:00:00', 129, 4),
(24, 703, '2020-06-09 00:00:00', 130, 4),
(25, 1728, '2020-06-09 00:00:00', 131, 4),
(26, 300, '2020-06-09 00:00:00', 144, 4),
(27, 300, '2020-06-09 00:00:00', 145, 4),
(28, 520, '2020-06-09 00:00:00', 146, 4),
(29, 520, '2020-06-09 00:00:00', 147, 4),
(30, 5184, '2020-06-09 00:00:00', 171, 4),
(31, 10368, '2020-06-09 00:00:00', 173, 4),
(32, 10368, '2020-06-09 00:00:00', 174, 4),
(33, 1728, '2020-06-09 00:00:00', 175, 4),
(34, 3300, '2020-06-09 00:00:00', 176, 4),
(35, 1030, '2020-06-10 00:00:00', 37, 5),
(36, 260, '2020-06-10 00:00:00', 122, 5),
(37, 260, '2020-06-10 00:00:00', 123, 5),
(38, 260, '2020-06-10 00:00:00', 124, 5),
(39, 260, '2020-06-10 00:00:00', 125, 5),
(40, 1100, '2020-06-10 00:00:00', 176, 5),
(41, 2000, '2020-06-10 00:00:00', 32, 6),
(42, 5184, '2020-06-10 00:00:00', 63, 6),
(43, 350, '2020-06-10 00:00:00', 104, 6),
(44, 700, '2020-06-10 00:00:00', 105, 6),
(45, 700, '2020-06-10 00:00:00', 106, 6),
(46, 350, '2020-06-10 00:00:00', 107, 6),
(47, 200, '2020-06-10 00:00:00', 144, 6),
(48, 400, '2020-06-10 00:00:00', 145, 6),
(49, 360, '2020-06-10 00:00:00', 146, 6),
(50, 200, '2020-06-10 00:00:00', 147, 6),
(51, 1700, '2020-06-10 00:00:00', 176, 6),
(52, 75, '2020-06-11 00:00:00', 38, 7),
(53, 75, '2020-06-11 00:00:00', 39, 7),
(54, 20, '2020-06-11 00:00:00', 31, 8),
(55, 20, '2020-06-11 00:00:00', 37, 8),
(56, 20, '2020-06-11 00:00:00', 129, 8),
(57, 5, '2020-06-11 00:00:00', 177, 8),
(58, 2, '2020-06-11 00:00:00', 178, 8),
(59, 1728, '2020-06-12 00:00:00', 53, 9),
(60, 3456, '2020-06-12 00:00:00', 63, 9),
(61, 1728, '2020-06-12 00:00:00', 169, 9),
(62, 1728, '2020-06-12 00:00:00', 170, 9),
(63, 1728, '2020-06-12 00:00:00', 171, 9),
(64, 3456, '2020-06-12 00:00:00', 173, 9),
(65, 310, '2020-06-12 00:00:00', 13, 10),
(66, 310, '2020-06-12 00:00:00', 14, 10),
(67, 620, '2020-06-12 00:00:00', 16, 10),
(68, 1900, '2020-06-12 00:00:00', 31, 10),
(69, 1900, '2020-06-12 00:00:00', 32, 10),
(70, 1900, '2020-06-12 00:00:00', 33, 10),
(71, 900, '2020-06-12 00:00:00', 37, 10),
(72, 3600, '2020-06-15 00:00:00', 34, 11),
(73, 3600, '2020-06-15 00:00:00', 41, 11),
(74, 1220, '2020-06-15 00:00:00', 31, 12),
(75, 1220, '2020-06-15 00:00:00', 41, 12),
(76, 210, '2020-06-15 00:00:00', 127, 12),
(77, 210, '2020-06-15 00:00:00', 128, 12),
(78, 420, '2020-06-15 00:00:00', 129, 12),
(79, 420, '2020-06-15 00:00:00', 130, 12),
(80, 90, '2020-06-16 00:00:00', 122, 13),
(81, 110, '2020-06-16 00:00:00', 123, 13),
(82, 60, '2020-06-16 00:00:00', 124, 13),
(83, 96, '2020-06-16 00:00:00', 125, 13),
(84, 930, '2020-06-18 00:00:00', 19, 14),
(85, 930, '2020-06-18 00:00:00', 20, 14),
(86, 930, '2020-06-18 00:00:00', 21, 14),
(87, 2500, '2020-06-18 00:00:00', 35, 14),
(88, 630, '2020-06-18 00:00:00', 57, 14),
(89, 630, '2020-06-18 00:00:00', 58, 14),
(90, 630, '2020-06-18 00:00:00', 59, 14),
(91, 630, '2020-06-18 00:00:00', 60, 14),
(92, 5184, '2020-06-18 00:00:00', 76, 14),
(93, 5184, '2020-06-18 00:00:00', 98, 14),
(94, 230, '2020-06-17 00:00:00', 60, 15),
(95, 180, '2020-06-17 00:00:00', 61, 15),
(96, 180, '2020-06-17 00:00:00', 62, 15),
(97, 310, '2020-06-17 00:00:00', 13, 16),
(98, 310, '2020-06-17 00:00:00', 14, 16),
(99, 620, '2020-06-17 00:00:00', 16, 16),
(100, 1900, '2020-06-17 00:00:00', 31, 16),
(101, 1900, '2020-06-17 00:00:00', 32, 16),
(102, 1900, '2020-06-17 00:00:00', 33, 16),
(103, 2900, '2020-06-17 00:00:00', 37, 16),
(104, 300, '2020-06-17 00:00:00', 13, 17),
(105, 300, '2020-06-17 00:00:00', 14, 17),
(106, 600, '2020-06-17 00:00:00', 16, 17),
(107, 1800, '2020-06-17 00:00:00', 31, 17),
(108, 1800, '2020-06-17 00:00:00', 32, 17),
(109, 1800, '2020-06-17 00:00:00', 33, 17),
(110, 1800, '2020-06-17 00:00:00', 37, 17),
(111, 300, '2020-06-17 00:00:00', 13, 18),
(112, 300, '2020-06-17 00:00:00', 14, 18),
(113, 600, '2020-06-17 00:00:00', 16, 18),
(114, 1800, '2020-06-17 00:00:00', 31, 18),
(115, 1800, '2020-06-17 00:00:00', 32, 18),
(116, 350, '2020-06-17 00:00:00', 127, 18),
(117, 350, '2020-06-17 00:00:00', 128, 18),
(118, 600, '2020-06-17 00:00:00', 129, 18),
(119, 500, '2020-06-17 00:00:00', 130, 18),
(120, 400, '2020-06-18 00:00:00', 14, 19),
(121, 200, '2020-06-18 00:00:00', 16, 19),
(122, 600, '2020-06-18 00:00:00', 17, 19),
(123, 1200, '2020-06-18 00:00:00', 31, 19),
(124, 1200, '2020-06-18 00:00:00', 37, 19),
(125, 200, '2020-06-18 00:00:00', 113, 19),
(126, 400, '2020-06-18 00:00:00', 128, 19),
(127, 400, '2020-06-18 00:00:00', 129, 19),
(139, 1600, '2020-06-19 00:00:00', 17, 21),
(140, 2000, '2020-06-19 00:00:00', 176, 21),
(141, 1200, '2020-06-19 00:00:00', 17, 22),
(142, 600, '2020-06-19 00:00:00', 24, 22),
(143, 600, '2020-06-19 00:00:00', 25, 22),
(144, 1200, '2020-06-19 00:00:00', 26, 22),
(145, 600, '2020-06-19 00:00:00', 27, 22),
(146, 2000, '2020-06-19 00:00:00', 176, 22),
(147, 1500, '2020-06-20 00:00:00', 31, 23),
(148, 250, '2020-06-20 00:00:00', 113, 23),
(149, 250, '2020-06-20 00:00:00', 128, 23),
(150, 500, '2020-06-20 00:00:00', 129, 23),
(151, 500, '2020-06-20 00:00:00', 35, 24),
(152, 230, '2020-06-20 00:00:00', 37, 24),
(153, 800, '2020-06-22 00:00:00', 176, 25),
(154, 200, '2020-06-22 00:00:00', 176, 26),
(155, 2100, '2020-06-22 00:00:00', 38, 27),
(156, 2100, '2020-06-22 00:00:00', 39, 27),
(157, 300, '2020-06-22 00:00:00', 13, 28),
(158, 250, '2020-06-22 00:00:00', 14, 28),
(159, 500, '2020-06-22 00:00:00', 16, 28),
(160, 1100, '2020-06-22 00:00:00', 24, 28),
(161, 1200, '2020-06-22 00:00:00', 25, 28),
(162, 2200, '2020-06-22 00:00:00', 26, 28),
(163, 900, '2020-06-22 00:00:00', 27, 28),
(164, 2400, '2020-06-22 00:00:00', 31, 28),
(165, 3900, '2020-06-22 00:00:00', 32, 28),
(166, 4104, '2020-06-22 00:00:00', 63, 28),
(167, 240, '2020-06-22 00:00:00', 127, 28),
(168, 260, '2020-06-22 00:00:00', 128, 28),
(169, 620, '2020-06-22 00:00:00', 129, 28),
(170, 850, '2020-06-22 00:00:00', 130, 28),
(171, 17280, '2020-06-22 00:00:00', 131, 28),
(172, 3456, '2020-06-22 00:00:00', 169, 28),
(173, 3456, '2020-06-22 00:00:00', 173, 28),
(174, 5184, '2020-06-22 00:00:00', 175, 28);

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
(4, '2020-06-09 00:00:00', 1),
(5, '2020-06-10 00:00:00', 3),
(6, '2020-06-10 00:00:00', 2),
(7, '2020-06-11 00:00:00', 9),
(8, '2020-06-11 00:00:00', 4),
(9, '2020-06-12 00:00:00', 3),
(10, '2020-06-12 00:00:00', 1),
(11, '2020-06-15 00:00:00', 11),
(12, '2020-06-15 00:00:00', 3),
(13, '2020-06-16 00:00:00', 3),
(14, '2020-06-18 00:00:00', 5),
(15, '2020-06-17 00:00:00', 8),
(16, '2020-06-17 00:00:00', 1),
(17, '2020-06-17 00:00:00', 2),
(18, '2020-06-17 00:00:00', 2),
(19, '2020-06-18 00:00:00', 3),
(21, '2020-06-19 00:00:00', 1),
(22, '2020-06-19 00:00:00', 2),
(23, '2020-06-20 00:00:00', 3),
(24, '2020-06-20 00:00:00', 8),
(25, '2020-06-22 00:00:00', 12),
(26, '2020-06-22 00:00:00', 13),
(27, '2020-06-22 00:00:00', 10),
(28, '2020-06-22 00:00:00', 1);

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
(12, 13, 3356),
(13, 14, 45),
(14, 15, 3),
(15, 16, 240),
(16, 17, 700),
(17, 18, 1620),
(18, 19, 3165),
(19, 20, 3370),
(20, 21, 761),
(21, 22, 100),
(22, 23, 32600),
(23, 24, 761),
(24, 25, 2830),
(25, 26, 730),
(26, 27, 310),
(27, 28, 3030),
(28, 29, 4000),
(29, 30, 3985),
(30, 31, 22060),
(31, 32, 29500),
(32, 33, 890),
(33, 34, 2000),
(34, 35, 150),
(35, 36, 500),
(36, 37, 100),
(37, 38, 2237),
(38, 39, 2237),
(40, 41, 6980),
(41, 42, 3660),
(42, 43, 3200),
(43, 44, 1670),
(44, 45, 6122),
(45, 46, 4060),
(46, 47, 4600),
(47, 48, 4600),
(48, 49, 100),
(49, 50, 2500),
(50, 51, 7027),
(51, 52, 6720),
(52, 53, 60408),
(53, 54, 8608),
(54, 55, 0),
(55, 56, 3920),
(56, 57, 3134),
(57, 58, 3270),
(58, 59, 1888),
(59, 60, 5155),
(60, 61, 8503),
(61, 62, 7060),
(62, 63, 5184),
(63, 64, 0),
(64, 65, 0),
(65, 66, 500),
(66, 67, 1400),
(67, 68, 1380),
(68, 69, 500),
(69, 70, 1672),
(70, 71, 1250),
(71, 72, 1324),
(72, 73, 1348),
(73, 74, 0),
(74, 75, 0),
(75, 76, 30240),
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
(90, 91, 31104),
(91, 92, 17496),
(92, 93, 1584),
(93, 94, 648),
(94, 95, 34848),
(95, 96, 36504),
(96, 97, 22464),
(97, 98, 49968),
(98, 99, 2160),
(99, 100, 2037),
(100, 101, 2035),
(101, 102, 4000),
(102, 103, 4000),
(103, 104, 1050),
(104, 105, 1215),
(105, 106, 3512),
(106, 107, 1368),
(107, 108, 1800),
(108, 109, 2275),
(109, 110, 2587),
(110, 111, 5270),
(111, 112, 2980),
(112, 113, 2032),
(113, 114, 700),
(114, 115, 2000),
(115, 116, 1200),
(116, 117, 800),
(117, 118, 350),
(118, 119, 600),
(119, 120, 750),
(120, 121, 400),
(121, 122, 33),
(122, 123, 608),
(123, 124, 594),
(124, 125, 146),
(125, 126, -166),
(126, 127, 2246),
(127, 128, 630),
(128, 129, 2778),
(129, 130, 2150),
(130, 131, 13824),
(131, 132, 4800),
(132, 133, 1000),
(133, 134, 1000),
(134, 135, 2000),
(135, 136, 2000),
(136, 137, 2050),
(137, 138, 800),
(138, 139, 900),
(139, 140, 1800),
(140, 141, 1500),
(141, 142, 1500),
(142, 143, 25000),
(143, 144, 1100),
(144, 145, 900),
(145, 146, 2600),
(146, 147, 2654),
(147, 148, 280),
(148, 149, 570),
(149, 150, 500),
(150, 151, 200),
(151, 152, 2020),
(152, 153, 2300),
(153, 154, 2450),
(154, 155, 924),
(155, 156, 380),
(156, 157, 714),
(157, 158, 840),
(158, 159, 1050),
(159, 160, 2100),
(160, 161, 2100),
(161, 162, 1050),
(162, 163, 1100),
(163, 164, 450),
(164, 165, 600),
(165, 166, 700),
(166, 167, 450),
(167, 168, 0),
(168, 169, 39240),
(169, 170, 37224),
(170, 171, 29952),
(171, 172, 36936),
(172, 173, 21672),
(173, 174, 25920),
(174, 175, 37728),
(175, 176, 0),
(176, 177, 0),
(177, 178, 0);

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
-- Indeks untuk tabel `sku_in`
--
ALTER TABLE `sku_in`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `sku_in_group`
--
ALTER TABLE `sku_in_group`
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
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- AUTO_INCREMENT untuk tabel `producer`
--
ALTER TABLE `producer`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT untuk tabel `product`
--
ALTER TABLE `product`
  MODIFY `id` int(200) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=187;

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
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=179;

--
-- AUTO_INCREMENT untuk tabel `sku_in`
--
ALTER TABLE `sku_in`
  MODIFY `id` int(200) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `sku_in_group`
--
ALTER TABLE `sku_in_group`
  MODIFY `id` int(200) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `sku_out`
--
ALTER TABLE `sku_out`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=175;

--
-- AUTO_INCREMENT untuk tabel `sku_out_group`
--
ALTER TABLE `sku_out_group`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=29;

--
-- AUTO_INCREMENT untuk tabel `sku_stock`
--
ALTER TABLE `sku_stock`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=178;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
