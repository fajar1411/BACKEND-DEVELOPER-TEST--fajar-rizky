-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 03 Mar 2023 pada 15.35
-- Versi server: 10.4.27-MariaDB
-- Versi PHP: 8.1.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `form`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `customers`
--

CREATE TABLE `customers` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` char(50) NOT NULL,
  `dob_date` longtext NOT NULL,
  `phonenum` varchar(20) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` longtext NOT NULL,
  `nationalities_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `customers`
--

INSERT INTO `customers` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `dob_date`, `phonenum`, `email`, `password`, `nationalities_id`) VALUES
(1, '2023-03-03 14:12:51.294', '2023-03-03 14:12:51.294', NULL, 'adrian', '28-09-2023', '0858884484', 'adrians1234@gmail.com', '$2a$10$MEHgkbcrucToaw93PvHT.OOHFMmePIHagpuCSz34y1cz6TRkZlpPi', 2),
(2, '2023-03-03 14:16:48.665', '2023-03-03 14:18:48.167', NULL, 'fajar rizky', '14-11-1997', '0858884484', 'frizky861@gmail.com', '$2a$10$jsc5MBjZATXT8g4PZGMfhuvSNysrFyH7Ym8DRijDnLdwpfQPYoo7i', 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `families`
--

CREATE TABLE `families` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `relation` varchar(50) NOT NULL,
  `name_relation` varchar(50) NOT NULL,
  `dob` varchar(50) NOT NULL,
  `customer_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `families`
--

INSERT INTO `families` (`id`, `created_at`, `updated_at`, `deleted_at`, `relation`, `name_relation`, `dob`, `customer_id`) VALUES
(2, '2023-03-03 14:20:05.629', '2023-03-03 14:32:07.946', NULL, 'abang', 'Rendi ', '05-06-1987', 2),
(3, '2023-03-03 14:20:34.405', '2023-03-03 14:20:34.405', '2023-03-03 14:27:10.324', 'abang', 'rendi dwinanto', '05-06-1987', 2),
(4, '2023-03-03 14:26:06.729', '2023-03-03 14:26:06.729', NULL, 'abang', 'faqih ', '05-06-1987', 2);

-- --------------------------------------------------------

--
-- Struktur dari tabel `nationalities`
--

CREATE TABLE `nationalities` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `nationality_name` varchar(50) NOT NULL,
  `nationality_code` char(2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `nationalities`
--

INSERT INTO `nationalities` (`id`, `created_at`, `updated_at`, `deleted_at`, `nationality_name`, `nationality_code`) VALUES
(1, '2023-03-03 11:48:59.955', '2023-03-03 11:48:59.955', NULL, 'indonesia', 'ID'),
(2, '2023-03-03 14:12:17.672', '2023-03-03 14:12:17.672', NULL, 'england', 'EN'),
(3, '2023-03-03 14:12:39.828', '2023-03-03 14:12:39.828', NULL, 'argentina', 'AR');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `idx_customers_deleted_at` (`deleted_at`),
  ADD KEY `fk_customers_nationalities` (`nationalities_id`);

--
-- Indeks untuk tabel `families`
--
ALTER TABLE `families`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_families_name_relation` (`name_relation`),
  ADD KEY `idx_families_deleted_at` (`deleted_at`),
  ADD KEY `fk_customers_families` (`customer_id`);

--
-- Indeks untuk tabel `nationalities`
--
ALTER TABLE `nationalities`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_nationalities_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `customers`
--
ALTER TABLE `customers`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `families`
--
ALTER TABLE `families`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `nationalities`
--
ALTER TABLE `nationalities`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `customers`
--
ALTER TABLE `customers`
  ADD CONSTRAINT `fk_customers_nationalities` FOREIGN KEY (`nationalities_id`) REFERENCES `nationalities` (`id`);

--
-- Ketidakleluasaan untuk tabel `families`
--
ALTER TABLE `families`
  ADD CONSTRAINT `fk_customers_families` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
