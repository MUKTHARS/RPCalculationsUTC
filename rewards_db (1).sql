-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 24, 2025 at 11:15 AM
-- Server version: 9.1.0
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `rewards_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `config`
--

CREATE TABLE `config` (
  `id` int NOT NULL,
  `config_key` varchar(50) NOT NULL,
  `config_value` text NOT NULL,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `config`
--

INSERT INTO `config` (`id`, `config_key`, `config_value`, `updated_at`) VALUES
(369, 'RedemptionRatioFirstYear', '15', '2025-05-22 05:29:00'),
(370, 'SubjectsCountFirstYear', '6', '2025-05-20 09:00:03'),
(371, 'MaxPointsPerSubjectFirstYear', '375', '2025-05-22 05:29:00'),
(372, 'RedemptionVarianceFirstYear', '[1,2,4,6]', '2025-05-20 09:00:03'),
(373, 'SectionPointsFirstYear', '[75,90,120,90]', '2025-05-22 05:29:00'),
(374, 'SectionMaxMarksFirstYear', '[5,3,2,1]', '2025-05-20 09:00:03'),
(375, 'PointsPerMarkFirstYear', '[15,30,60,90]', '2025-05-22 05:29:00'),
(376, 'MinPointsForHalfFirstYear', '[7.5,15,30,45]', '2025-05-22 05:29:00'),
(377, 'RedemptionRatioSecondYear', '50', '2025-05-20 09:00:12'),
(378, 'SubjectsCountSecondYear', '6', '2025-05-20 09:00:12'),
(379, 'MaxPointsPerSubjectSecondYear', '1250', '2025-05-20 09:00:12'),
(380, 'RedemptionVarianceSecondYear', '[1,2,4,6]', '2025-05-20 09:00:12'),
(381, 'SectionPointsSecondYear', '[250,300,400,300]', '2025-05-20 09:00:12'),
(382, 'SectionMaxMarksSecondYear', '[5,3,2,1]', '2025-05-20 09:00:12'),
(383, 'PointsPerMarkSecondYear', '[50,100,200,300]', '2025-05-20 09:00:12'),
(384, 'MinPointsForHalfSecondYear', '[25,50,100,150]', '2025-05-20 09:00:12'),
(385, 'RedemptionRatioThirdYear', '90', '2025-05-20 09:00:24'),
(386, 'SubjectsCountThirdYear', '6', '2025-05-20 09:00:24'),
(387, 'MaxPointsPerSubjectThirdYear', '2250', '2025-05-20 09:00:24'),
(388, 'RedemptionVarianceThirdYear', '[1,2,4,6]', '2025-05-20 09:00:24'),
(389, 'SectionPointsThirdYear', '[450,540,720,540]', '2025-05-20 09:00:24'),
(390, 'SectionMaxMarksThirdYear', '[5,3,2,1]', '2025-05-20 09:00:24'),
(391, 'PointsPerMarkThirdYear', '[90,180,360,540]', '2025-05-20 09:00:24'),
(392, 'MinPointsForHalfThirdYear', '[45,90,180,270]', '2025-05-20 09:00:24'),
(393, 'RedemptionRatioFinalYear', '120', '2025-05-20 09:00:32'),
(394, 'SubjectsCountFinalYear', '6', '2025-05-20 09:00:32'),
(395, 'MaxPointsPerSubjectFinalYear', '3000', '2025-05-20 09:00:32'),
(396, 'RedemptionVarianceFinalYear', '[1,2,4,6]', '2025-05-20 09:00:32'),
(397, 'SectionPointsFinalYear', '[600,720,960,720]', '2025-05-20 09:00:32'),
(398, 'SectionMaxMarksFinalYear', '[5,3,2,1]', '2025-05-20 09:00:32'),
(399, 'PointsPerMarkFinalYear', '[120,240,480,720]', '2025-05-20 09:00:32'),
(400, 'MinPointsForHalfFinalYear', '[60,120,240,360]', '2025-05-20 09:00:32'),
(401, 'PointsPerMark', '[15,30,60,90]', '2025-05-22 05:29:00'),
(402, 'MinPointsForHalf', '[7.5,15,30,45]', '2025-05-22 05:29:00'),
(403, 'RedemptionRatio', '15', '2025-05-22 05:29:00'),
(404, 'SubjectsCount', '6', '2025-05-20 10:29:32'),
(405, 'MaxPointsPerSubject', '375', '2025-05-22 05:29:00'),
(406, 'RedemptionVariance', '[1,2,4,6]', '2025-05-20 10:29:32'),
(407, 'SectionPoints', '[75,90,120,90]', '2025-05-22 05:29:00'),
(408, 'SectionMaxMarks', '[5,3,2,1]', '2025-05-20 10:29:32'),
(465, 'RedemptionRatioFirstYearGENERAL', '10', '2025-05-23 04:18:34'),
(466, 'SubjectsCountFirstYearGENERAL', '6', '2025-05-23 04:18:34'),
(467, 'MaxPointsPerSubjectFirstYearGENERAL', '250', '2025-05-23 04:18:34'),
(468, 'RedemptionVarianceFirstYearGENERAL', '[1,2,4,6]', '2025-05-23 04:18:34'),
(469, 'SectionPointsFirstYearGENERAL', '[50,60,80,60]', '2025-05-23 04:18:34'),
(470, 'SectionMaxMarksFirstYearGENERAL', '[5,3,2,1]', '2025-05-23 04:18:34'),
(471, 'PointsPerMarkFirstYearGENERAL', '[10,20,40,60]', '2025-05-23 04:18:34'),
(472, 'MinPointsForHalfFirstYearGENERAL', '[5,10,20,30]', '2025-05-23 04:18:34'),
(473, 'RedemptionRatioSecondYearGENERAL', '50', '2025-05-23 04:18:34'),
(474, 'SubjectsCountSecondYearGENERAL', '6', '2025-05-23 04:18:34'),
(475, 'MaxPointsPerSubjectSecondYearGENERAL', '1250', '2025-05-23 04:18:34'),
(476, 'RedemptionVarianceSecondYearGENERAL', '[1,2,4,6]', '2025-05-23 04:18:34'),
(477, 'SectionPointsSecondYearGENERAL', '[250,300,400,300]', '2025-05-23 04:18:34'),
(478, 'SectionMaxMarksSecondYearGENERAL', '[5,3,2,1]', '2025-05-23 04:18:34'),
(479, 'PointsPerMarkSecondYearGENERAL', '[50,100,200,300]', '2025-05-23 04:18:34'),
(480, 'MinPointsForHalfSecondYearGENERAL', '[25,50,100,150]', '2025-05-23 04:18:34'),
(481, 'RedemptionRatioThirdYearGENERAL', '90', '2025-05-23 04:18:34'),
(482, 'SubjectsCountThirdYearGENERAL', '6', '2025-05-23 04:18:34'),
(483, 'MaxPointsPerSubjectThirdYearGENERAL', '2250', '2025-05-23 04:18:34'),
(484, 'RedemptionVarianceThirdYearGENERAL', '[1,2,4,6]', '2025-05-23 04:18:34'),
(485, 'SectionPointsThirdYearGENERAL', '[450,540,720,540]', '2025-05-23 04:18:34'),
(486, 'SectionMaxMarksThirdYearGENERAL', '[5,3,2,1]', '2025-05-23 04:18:34'),
(487, 'PointsPerMarkThirdYearGENERAL', '[90,180,360,540]', '2025-05-23 04:18:34'),
(488, 'MinPointsForHalfThirdYearGENERAL', '[45,90,180,270]', '2025-05-23 04:18:34'),
(489, 'RedemptionRatioFinalYearGENERAL', '120', '2025-05-23 04:18:34'),
(490, 'SubjectsCountFinalYearGENERAL', '6', '2025-05-23 04:18:34'),
(491, 'MaxPointsPerSubjectFinalYearGENERAL', '3000', '2025-05-23 04:18:34'),
(492, 'RedemptionVarianceFinalYearGENERAL', '[1,2,4,6]', '2025-05-23 04:18:34'),
(493, 'SectionPointsFinalYearGENERAL', '[600,720,960,720]', '2025-05-23 04:18:34'),
(494, 'SectionMaxMarksFinalYearGENERAL', '[5,3,2,1]', '2025-05-23 04:18:34'),
(495, 'PointsPerMarkFinalYearGENERAL', '[120,240,480,720]', '2025-05-23 04:18:34'),
(496, 'MinPointsForHalfFinalYearGENERAL', '[60,120,240,360]', '2025-05-23 04:18:34'),
(497, 'RedemptionRatioFirstYearCSE', '15', '2025-05-23 04:18:34'),
(498, 'SubjectsCountFirstYearCSE', '6', '2025-05-23 04:18:34'),
(499, 'MaxPointsPerSubjectFirstYearCSE', '375', '2025-05-23 04:18:34'),
(500, 'RedemptionVarianceFirstYearCSE', '[1,2,4,6]', '2025-05-23 04:18:34'),
(501, 'SectionPointsFirstYearCSE', '[75,90,120,90]', '2025-05-23 04:18:34'),
(502, 'SectionMaxMarksFirstYearCSE', '[5,3,2,1]', '2025-05-23 04:18:34'),
(503, 'PointsPerMarkFirstYearCSE', '[15,30,60,90]', '2025-05-23 04:18:34'),
(504, 'MinPointsForHalfFirstYearCSE', '[7.5,15,30,45]', '2025-05-23 04:18:34'),
(505, 'RedemptionRatioSecondYearCSE', '55', '2025-05-23 04:18:34'),
(506, 'SubjectsCountSecondYearCSE', '6', '2025-05-23 04:18:34'),
(507, 'MaxPointsPerSubjectSecondYearCSE', '1375', '2025-05-23 04:18:34'),
(508, 'RedemptionVarianceSecondYearCSE', '[1,2,4,6]', '2025-05-23 04:18:34'),
(509, 'SectionPointsSecondYearCSE', '[275,330,440,330]', '2025-05-23 04:18:34'),
(510, 'SectionMaxMarksSecondYearCSE', '[5,3,2,1]', '2025-05-23 04:18:34'),
(511, 'PointsPerMarkSecondYearCSE', '[55,110,220,330]', '2025-05-23 04:18:34'),
(512, 'MinPointsForHalfSecondYearCSE', '[27.5,55,110,165]', '2025-05-23 04:18:34'),
(513, 'RedemptionRatioThirdYearCSE', '95', '2025-05-23 04:18:34'),
(514, 'SubjectsCountThirdYearCSE', '6', '2025-05-23 04:18:34'),
(515, 'MaxPointsPerSubjectThirdYearCSE', '2375', '2025-05-23 04:18:34'),
(516, 'RedemptionVarianceThirdYearCSE', '[1,2,4,6]', '2025-05-23 04:18:34'),
(517, 'SectionPointsThirdYearCSE', '[475,570,760,570]', '2025-05-23 04:18:34'),
(518, 'SectionMaxMarksThirdYearCSE', '[5,3,2,1]', '2025-05-23 04:18:34'),
(519, 'PointsPerMarkThirdYearCSE', '[95,190,380,570]', '2025-05-23 04:18:34'),
(520, 'MinPointsForHalfThirdYearCSE', '[47.5,95,190,285]', '2025-05-23 04:18:34'),
(521, 'RedemptionRatioFinalYearCSE', '55', '2025-05-23 09:15:03'),
(522, 'SubjectsCountFinalYearCSE', '6', '2025-05-23 04:18:34'),
(523, 'MaxPointsPerSubjectFinalYearCSE', '1375', '2025-05-23 09:15:03'),
(524, 'RedemptionVarianceFinalYearCSE', '[1,2,4,6]', '2025-05-23 04:18:34'),
(525, 'SectionPointsFinalYearCSE', '[275,330,440,330]', '2025-05-23 09:15:03'),
(526, 'SectionMaxMarksFinalYearCSE', '[5,3,2,1]', '2025-05-23 04:18:34'),
(527, 'PointsPerMarkFinalYearCSE', '[55,110,220,330]', '2025-05-23 09:15:03'),
(528, 'MinPointsForHalfFinalYearCSE', '[27.5,55,110,165]', '2025-05-23 09:15:03'),
(529, 'RedemptionRatioFirstYearECE', '10', '2025-05-23 04:18:34'),
(530, 'SubjectsCountFirstYearECE', '6', '2025-05-23 04:18:34'),
(531, 'MaxPointsPerSubjectFirstYearECE', '250', '2025-05-23 04:18:34'),
(532, 'RedemptionVarianceFirstYearECE', '[1,2,4,6]', '2025-05-23 04:18:34'),
(533, 'SectionPointsFirstYearECE', '[50,60,80,60]', '2025-05-23 04:18:34'),
(534, 'SectionMaxMarksFirstYearECE', '[5,3,2,1]', '2025-05-23 04:18:34'),
(535, 'PointsPerMarkFirstYearECE', '[10,20,40,60]', '2025-05-23 04:18:34'),
(536, 'MinPointsForHalfFirstYearECE', '[5,10,20,30]', '2025-05-23 04:18:34'),
(537, 'RedemptionRatioSecondYearECE', '50', '2025-05-23 04:18:34'),
(538, 'SubjectsCountSecondYearECE', '6', '2025-05-23 04:18:34'),
(539, 'MaxPointsPerSubjectSecondYearECE', '1250', '2025-05-23 04:18:34'),
(540, 'RedemptionVarianceSecondYearECE', '[1,2,4,6]', '2025-05-23 04:18:34'),
(541, 'SectionPointsSecondYearECE', '[250,300,400,300]', '2025-05-23 04:18:34'),
(542, 'SectionMaxMarksSecondYearECE', '[5,3,2,1]', '2025-05-23 04:18:34'),
(543, 'PointsPerMarkSecondYearECE', '[50,100,200,300]', '2025-05-23 04:18:34'),
(544, 'MinPointsForHalfSecondYearECE', '[25,50,100,150]', '2025-05-23 04:18:34'),
(545, 'RedemptionRatioThirdYearECE', '90', '2025-05-23 04:18:34'),
(546, 'SubjectsCountThirdYearECE', '6', '2025-05-23 04:18:34'),
(547, 'MaxPointsPerSubjectThirdYearECE', '2250', '2025-05-23 04:18:34'),
(548, 'RedemptionVarianceThirdYearECE', '[1,2,4,6]', '2025-05-23 04:18:34'),
(549, 'SectionPointsThirdYearECE', '[450,540,720,540]', '2025-05-23 04:18:34'),
(550, 'SectionMaxMarksThirdYearECE', '[5,3,2,1]', '2025-05-23 04:18:34'),
(551, 'PointsPerMarkThirdYearECE', '[90,180,360,540]', '2025-05-23 04:18:34'),
(552, 'MinPointsForHalfThirdYearECE', '[45,90,180,270]', '2025-05-23 04:18:34'),
(553, 'RedemptionRatioFinalYearECE', '120', '2025-05-23 04:18:34'),
(554, 'SubjectsCountFinalYearECE', '6', '2025-05-23 04:18:34'),
(555, 'MaxPointsPerSubjectFinalYearECE', '3000', '2025-05-23 04:18:34'),
(556, 'RedemptionVarianceFinalYearECE', '[1,2,4,6]', '2025-05-23 04:18:34'),
(557, 'SectionPointsFinalYearECE', '[600,720,960,720]', '2025-05-23 04:18:34'),
(558, 'SectionMaxMarksFinalYearECE', '[5,3,2,1]', '2025-05-23 04:18:34'),
(559, 'PointsPerMarkFinalYearECE', '[120,240,480,720]', '2025-05-23 04:18:34'),
(560, 'MinPointsForHalfFinalYearECE', '[60,120,240,360]', '2025-05-23 04:18:34'),
(561, 'RedemptionRatioFirstYearEEE', '10', '2025-05-23 04:18:34'),
(562, 'SubjectsCountFirstYearEEE', '6', '2025-05-23 04:18:34'),
(563, 'MaxPointsPerSubjectFirstYearEEE', '250', '2025-05-23 04:18:34'),
(564, 'RedemptionVarianceFirstYearEEE', '[1,2,4,6]', '2025-05-23 04:18:34'),
(565, 'SectionPointsFirstYearEEE', '[50,60,80,60]', '2025-05-23 04:18:34'),
(566, 'SectionMaxMarksFirstYearEEE', '[5,3,2,1]', '2025-05-23 04:18:34'),
(567, 'PointsPerMarkFirstYearEEE', '[10,20,40,60]', '2025-05-23 04:18:34'),
(568, 'MinPointsForHalfFirstYearEEE', '[5,10,20,30]', '2025-05-23 04:18:34'),
(569, 'RedemptionRatioSecondYearEEE', '30', '2025-05-23 06:12:36'),
(570, 'SubjectsCountSecondYearEEE', '6', '2025-05-23 04:18:34'),
(571, 'MaxPointsPerSubjectSecondYearEEE', '750', '2025-05-23 06:12:36'),
(572, 'RedemptionVarianceSecondYearEEE', '[1,2,4,6]', '2025-05-23 04:18:34'),
(573, 'SectionPointsSecondYearEEE', '[150,180,240,180]', '2025-05-23 06:12:36'),
(574, 'SectionMaxMarksSecondYearEEE', '[5,3,2,1]', '2025-05-23 04:18:34'),
(575, 'PointsPerMarkSecondYearEEE', '[30,60,120,180]', '2025-05-23 06:12:36'),
(576, 'MinPointsForHalfSecondYearEEE', '[15,30,60,90]', '2025-05-23 06:12:36'),
(577, 'RedemptionRatioThirdYearEEE', '90', '2025-05-23 04:18:34'),
(578, 'SubjectsCountThirdYearEEE', '6', '2025-05-23 04:18:34'),
(579, 'MaxPointsPerSubjectThirdYearEEE', '2250', '2025-05-23 04:18:34'),
(580, 'RedemptionVarianceThirdYearEEE', '[1,2,4,6]', '2025-05-23 04:18:34'),
(581, 'SectionPointsThirdYearEEE', '[450,540,720,540]', '2025-05-23 04:18:34'),
(582, 'SectionMaxMarksThirdYearEEE', '[5,3,2,1]', '2025-05-23 04:18:34'),
(583, 'PointsPerMarkThirdYearEEE', '[90,180,360,540]', '2025-05-23 04:18:34'),
(584, 'MinPointsForHalfThirdYearEEE', '[45,90,180,270]', '2025-05-23 04:18:34'),
(585, 'RedemptionRatioFinalYearEEE', '120', '2025-05-23 04:18:34'),
(586, 'SubjectsCountFinalYearEEE', '6', '2025-05-23 04:18:34'),
(587, 'MaxPointsPerSubjectFinalYearEEE', '3000', '2025-05-23 04:18:34'),
(588, 'RedemptionVarianceFinalYearEEE', '[1,2,4,6]', '2025-05-23 04:18:34'),
(589, 'SectionPointsFinalYearEEE', '[600,720,960,720]', '2025-05-23 04:18:34'),
(590, 'SectionMaxMarksFinalYearEEE', '[5,3,2,1]', '2025-05-23 04:18:34'),
(591, 'PointsPerMarkFinalYearEEE', '[120,240,480,720]', '2025-05-23 04:18:34'),
(592, 'MinPointsForHalfFinalYearEEE', '[60,120,240,360]', '2025-05-23 04:18:34'),
(593, 'RedemptionRatioFirstYearMECH', '10', '2025-05-23 04:18:34'),
(594, 'SubjectsCountFirstYearMECH', '6', '2025-05-23 04:18:34'),
(595, 'MaxPointsPerSubjectFirstYearMECH', '250', '2025-05-23 04:18:34'),
(596, 'RedemptionVarianceFirstYearMECH', '[1,2,4,6]', '2025-05-23 04:18:34'),
(597, 'SectionPointsFirstYearMECH', '[50,60,80,60]', '2025-05-23 04:18:34'),
(598, 'SectionMaxMarksFirstYearMECH', '[5,3,2,1]', '2025-05-23 04:18:34'),
(599, 'PointsPerMarkFirstYearMECH', '[10,20,40,60]', '2025-05-23 04:18:34'),
(600, 'MinPointsForHalfFirstYearMECH', '[5,10,20,30]', '2025-05-23 04:18:34'),
(601, 'RedemptionRatioSecondYearMECH', '50', '2025-05-23 04:18:34'),
(602, 'SubjectsCountSecondYearMECH', '6', '2025-05-23 04:18:34'),
(603, 'MaxPointsPerSubjectSecondYearMECH', '1250', '2025-05-23 04:18:34'),
(604, 'RedemptionVarianceSecondYearMECH', '[1,2,4,6]', '2025-05-23 04:18:34'),
(605, 'SectionPointsSecondYearMECH', '[250,300,400,300]', '2025-05-23 04:18:34'),
(606, 'SectionMaxMarksSecondYearMECH', '[5,3,2,1]', '2025-05-23 04:18:34'),
(607, 'PointsPerMarkSecondYearMECH', '[50,100,200,300]', '2025-05-23 04:18:34'),
(608, 'MinPointsForHalfSecondYearMECH', '[25,50,100,150]', '2025-05-23 04:18:34'),
(609, 'RedemptionRatioThirdYearMECH', '90', '2025-05-23 04:18:34'),
(610, 'SubjectsCountThirdYearMECH', '6', '2025-05-23 04:18:34'),
(611, 'MaxPointsPerSubjectThirdYearMECH', '2250', '2025-05-23 04:18:34'),
(612, 'RedemptionVarianceThirdYearMECH', '[1,2,4,6]', '2025-05-23 04:18:34'),
(613, 'SectionPointsThirdYearMECH', '[450,540,720,540]', '2025-05-23 04:18:34'),
(614, 'SectionMaxMarksThirdYearMECH', '[5,3,2,1]', '2025-05-23 04:18:34'),
(615, 'PointsPerMarkThirdYearMECH', '[90,180,360,540]', '2025-05-23 04:18:34'),
(616, 'MinPointsForHalfThirdYearMECH', '[45,90,180,270]', '2025-05-23 04:18:34'),
(617, 'RedemptionRatioFinalYearMECH', '120', '2025-05-23 04:18:34'),
(618, 'SubjectsCountFinalYearMECH', '6', '2025-05-23 04:18:34'),
(619, 'MaxPointsPerSubjectFinalYearMECH', '3000', '2025-05-23 04:18:34'),
(620, 'RedemptionVarianceFinalYearMECH', '[1,2,4,6]', '2025-05-23 04:18:34'),
(621, 'SectionPointsFinalYearMECH', '[600,720,960,720]', '2025-05-23 04:18:34'),
(622, 'SectionMaxMarksFinalYearMECH', '[5,3,2,1]', '2025-05-23 04:18:34'),
(623, 'PointsPerMarkFinalYearMECH', '[120,240,480,720]', '2025-05-23 04:18:34'),
(624, 'MinPointsForHalfFinalYearMECH', '[60,120,240,360]', '2025-05-23 04:18:34'),
(673, 'RedemptionRatioFirstYearIT', '20', '2025-05-23 06:08:39'),
(674, 'SubjectsCountFirstYearIT', '6', '2025-05-23 06:08:39'),
(675, 'MaxPointsPerSubjectFirstYearIT', '500', '2025-05-23 06:08:39'),
(676, 'RedemptionVarianceFirstYearIT', '[1,2,4,6]', '2025-05-23 06:08:39'),
(677, 'SectionPointsFirstYearIT', '[100,120,160,120]', '2025-05-23 06:08:39'),
(678, 'SectionMaxMarksFirstYearIT', '[5,3,2,1]', '2025-05-23 06:08:39'),
(679, 'PointsPerMarkFirstYearIT', '[20,40,80,120]', '2025-05-23 06:08:39'),
(680, 'MinPointsForHalfFirstYearIT', '[10,20,40,60]', '2025-05-23 06:08:39');

-- --------------------------------------------------------

--
-- Table structure for table `config_year`
--

CREATE TABLE `config_year` (
  `year` int NOT NULL,
  `config_data` text NOT NULL,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `department` varchar(50) NOT NULL DEFAULT 'IT'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `config_year`
--

INSERT INTO `config_year` (`year`, `config_data`, `updated_at`, `department`) VALUES
(1, '{\"redemptionRatio\":15,\"subjectsCount\":6,\"maxPointsPerSubject\":375,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[75,90,120,90],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[15,30,60,90],\"minPointsForHalf\":[7.5,15,30,45],\"subjectNames\":null,\"department\":\"CSE\"}', '2025-05-23 05:42:39', 'CSE'),
(1, '{\"redemptionRatio\":10,\"subjectsCount\":6,\"maxPointsPerSubject\":250,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[50,60,80,60],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[10,20,40,60],\"minPointsForHalf\":[5,10,20,30],\"subjectNames\":null,\"department\":\"ECE\"}', '2025-05-23 05:42:03', 'ECE'),
(1, '{\"redemptionRatio\":10,\"subjectsCount\":6,\"maxPointsPerSubject\":250,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[50,60,80,60],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[10,20,40,60],\"minPointsForHalf\":[5,10,20,30],\"subjectNames\":null,\"department\":\"EEE\"}', '2025-05-23 05:41:54', 'EEE'),
(1, '{\"redemptionRatio\":10,\"subjectsCount\":6,\"maxPointsPerSubject\":250,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[50,60,80,60],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[10,20,40,60],\"minPointsForHalf\":[5,10,20,30],\"subjectNames\":null,\"department\":\"GENERAL\"}', '2025-05-23 04:18:34', 'GENERAL'),
(1, '{\"redemptionRatio\":20,\"subjectsCount\":6,\"maxPointsPerSubject\":500,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[100,120,160,120],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[20,40,80,120],\"minPointsForHalf\":[10,20,40,60],\"subjectNames\":null,\"department\":\"IT\"}', '2025-05-23 06:08:39', 'IT'),
(1, '{\"redemptionRatio\":10,\"subjectsCount\":6,\"maxPointsPerSubject\":250,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[50,60,80,60],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[10,20,40,60],\"minPointsForHalf\":[5,10,20,30],\"subjectNames\":null,\"department\":\"MECH\"}', '2025-05-23 04:18:34', 'MECH'),
(2, '{\"redemptionRatio\":55,\"subjectsCount\":6,\"maxPointsPerSubject\":1375,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[275,330,440,330],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[55,110,220,330],\"minPointsForHalf\":[27.5,55,110,165],\"subjectNames\":null,\"department\":\"CSE\"}', '2025-05-23 04:18:34', 'CSE'),
(2, '{\"redemptionRatio\":50,\"subjectsCount\":6,\"maxPointsPerSubject\":1250,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[250,300,400,300],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[50,100,200,300],\"minPointsForHalf\":[25,50,100,150],\"subjectNames\":null,\"department\":\"ECE\"}', '2025-05-23 04:18:34', 'ECE'),
(2, '{\"redemptionRatio\":30,\"subjectsCount\":6,\"maxPointsPerSubject\":750,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[150,180,240,180],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[30,60,120,180],\"minPointsForHalf\":[15,30,60,90],\"subjectNames\":null,\"department\":\"EEE\"}', '2025-05-23 06:12:36', 'EEE'),
(2, '{\"redemptionRatio\":50,\"subjectsCount\":6,\"maxPointsPerSubject\":1250,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[250,300,400,300],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[50,100,200,300],\"minPointsForHalf\":[25,50,100,150],\"subjectNames\":null,\"department\":\"GENERAL\"}', '2025-05-23 04:18:34', 'GENERAL'),
(2, '{\"redemptionRatio\":50,\"subjectsCount\":6,\"maxPointsPerSubject\":1250,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[250,300,400,300],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[50,100,200,300],\"minPointsForHalf\":[25,50,100,150],\"subjectNames\":null}', '2025-05-22 05:29:53', 'IT'),
(2, '{\"redemptionRatio\":50,\"subjectsCount\":6,\"maxPointsPerSubject\":1250,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[250,300,400,300],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[50,100,200,300],\"minPointsForHalf\":[25,50,100,150],\"subjectNames\":null,\"department\":\"MECH\"}', '2025-05-23 04:18:34', 'MECH'),
(3, '{\"redemptionRatio\":95,\"subjectsCount\":6,\"maxPointsPerSubject\":2375,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[475,570,760,570],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[95,190,380,570],\"minPointsForHalf\":[47.5,95,190,285],\"subjectNames\":null,\"department\":\"CSE\"}', '2025-05-23 04:18:34', 'CSE'),
(3, '{\"redemptionRatio\":90,\"subjectsCount\":6,\"maxPointsPerSubject\":2250,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[450,540,720,540],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[90,180,360,540],\"minPointsForHalf\":[45,90,180,270],\"subjectNames\":null,\"department\":\"ECE\"}', '2025-05-23 04:18:34', 'ECE'),
(3, '{\"redemptionRatio\":90,\"subjectsCount\":6,\"maxPointsPerSubject\":2250,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[450,540,720,540],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[90,180,360,540],\"minPointsForHalf\":[45,90,180,270],\"subjectNames\":null,\"department\":\"EEE\"}', '2025-05-23 04:18:34', 'EEE'),
(3, '{\"redemptionRatio\":90,\"subjectsCount\":6,\"maxPointsPerSubject\":2250,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[450,540,720,540],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[90,180,360,540],\"minPointsForHalf\":[45,90,180,270],\"subjectNames\":null,\"department\":\"GENERAL\"}', '2025-05-23 04:18:34', 'GENERAL'),
(3, '{\"redemptionRatio\":90,\"subjectsCount\":6,\"maxPointsPerSubject\":2250,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[450,540,720,540],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[90,180,360,540],\"minPointsForHalf\":[45,90,180,270],\"subjectNames\":null}', '2025-05-20 09:00:24', 'IT'),
(3, '{\"redemptionRatio\":90,\"subjectsCount\":6,\"maxPointsPerSubject\":2250,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[450,540,720,540],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[90,180,360,540],\"minPointsForHalf\":[45,90,180,270],\"subjectNames\":null,\"department\":\"MECH\"}', '2025-05-23 04:18:34', 'MECH'),
(4, '{\"redemptionRatio\":55,\"subjectsCount\":6,\"maxPointsPerSubject\":1375,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[275,330,440,330],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[55,110,220,330],\"minPointsForHalf\":[27.5,55,110,165],\"subjectNames\":null,\"department\":\"CSE\"}', '2025-05-23 09:15:03', 'CSE'),
(4, '{\"redemptionRatio\":120,\"subjectsCount\":6,\"maxPointsPerSubject\":3000,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[600,720,960,720],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[120,240,480,720],\"minPointsForHalf\":[60,120,240,360],\"subjectNames\":null,\"department\":\"ECE\"}', '2025-05-23 04:18:34', 'ECE'),
(4, '{\"redemptionRatio\":120,\"subjectsCount\":6,\"maxPointsPerSubject\":3000,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[600,720,960,720],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[120,240,480,720],\"minPointsForHalf\":[60,120,240,360],\"subjectNames\":null,\"department\":\"EEE\"}', '2025-05-23 04:18:34', 'EEE'),
(4, '{\"redemptionRatio\":120,\"subjectsCount\":6,\"maxPointsPerSubject\":3000,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[600,720,960,720],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[120,240,480,720],\"minPointsForHalf\":[60,120,240,360],\"subjectNames\":null,\"department\":\"GENERAL\"}', '2025-05-23 04:18:34', 'GENERAL'),
(4, '{\"redemptionRatio\":120,\"subjectsCount\":6,\"maxPointsPerSubject\":3000,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[600,720,960,720],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[120,240,480,720],\"minPointsForHalf\":[60,120,240,360],\"subjectNames\":null}', '2025-05-20 09:00:32', 'IT'),
(4, '{\"redemptionRatio\":120,\"subjectsCount\":6,\"maxPointsPerSubject\":3000,\"redemptionVariance\":[1,2,4,6],\"sectionPoints\":[600,720,960,720],\"sectionMaxMarks\":[5,3,2,1],\"pointsPerMark\":[120,240,480,720],\"minPointsForHalf\":[60,120,240,360],\"subjectNames\":null,\"department\":\"MECH\"}', '2025-05-23 04:18:34', 'MECH');

-- --------------------------------------------------------

--
-- Table structure for table `departments`
--

CREATE TABLE `departments` (
  `code` varchar(10) NOT NULL,
  `name` varchar(100) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `departments`
--

INSERT INTO `departments` (`code`, `name`, `created_at`) VALUES
('AERO', 'Aeronautical Engineering', '2025-05-23 10:29:39'),
('CIIVL', 'Civil Engineering', '2025-05-23 08:57:10'),
('CSE', 'Computer Science', '2025-05-23 08:56:26'),
('ECE', 'Electronics and Communication', '2025-05-23 08:56:26'),
('EEE', 'Electrical and Electronics', '2025-05-23 08:56:26'),
('IT', 'Information Technology', '2025-05-23 08:56:26'),
('MECH', 'Mechanical', '2025-05-23 08:56:26');

-- --------------------------------------------------------

--
-- Table structure for table `redemption_dates`
--

CREATE TABLE `redemption_dates` (
  `year` int NOT NULL,
  `department` varchar(50) NOT NULL,
  `start_date` datetime NOT NULL,
  `end_date` datetime NOT NULL,
  `is_extended` tinyint(1) DEFAULT '0',
  `extended_date` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `redemption_dates`
--

INSERT INTO `redemption_dates` (`year`, `department`, `start_date`, `end_date`, `is_extended`, `extended_date`) VALUES
(1, 'GENERAL', '2025-05-20 09:07:00', '2025-05-23 09:07:00', 0, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `reward_config`
--

CREATE TABLE `reward_config` (
  `id` int NOT NULL,
  `config_key` varchar(50) NOT NULL,
  `config_value` text NOT NULL,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `reward_config`
--

INSERT INTO `reward_config` (`id`, `config_key`, `config_value`, `updated_at`) VALUES
(1, 'RedemptionRatio', '10', '2025-05-12 09:45:16'),
(2, 'MaxPointsPerSubject', '250', '2025-05-12 09:45:16'),
(3, 'RedemptionVariance', '[1,2,4,6]', '2025-05-10 09:36:27'),
(4, 'SectionPoints', '[50,60,80,60]', '2025-05-12 09:45:16'),
(5, 'SectionMaxMarks', '[5,3,2,1]', '2025-05-08 03:08:30'),
(6, 'PointsPerMark', '[10,20,40,60]', '2025-05-12 09:45:16'),
(7, 'MinPointsForHalf', '[5,10,20,30]', '2025-05-12 09:45:16');

-- --------------------------------------------------------

--
-- Table structure for table `student_rewards`
--

CREATE TABLE `student_rewards` (
  `id` int NOT NULL,
  `student_id` varchar(50) NOT NULL,
  `student_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `student_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `department` varchar(100) DEFAULT NULL,
  `year` int NOT NULL DEFAULT '1',
  `points` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `total_marks` float NOT NULL DEFAULT '0',
  `points_remaining` int NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `student_rewards`
--

INSERT INTO `student_rewards` (`id`, `student_id`, `student_name`, `student_email`, `department`, `year`, `points`, `created_at`, `total_marks`, `points_remaining`) VALUES
(1, 'S001', 'John Smith', 'john.smith@university.edu', 'IT', 2025, 297, '2025-05-07 03:53:03', 19, 5),
(2, 'S002', 'Emily Johnson', 'emily.johnson@university.edu', 'CSE', 1, 1410, '2025-05-07 03:53:03', 55, 0),
(3, 'S003', 'Michael Williams', 'michael.williams@university.edu', 'CSD', 1, 1500, '2025-05-07 03:53:03', 66, 0),
(4, 'S004', 'Sarah Brown', 'sarah.brown@university.edu', 'CT', 1, 1900, '2025-05-07 03:53:03', 62, 10),
(5, 'S005', 'David Jones', 'david.jones@university.edu', 'ISE', 1, 500, '2025-05-07 03:53:03', 31.5, 5),
(6, 'S006', 'Jessica Garcia', 'jessica.garcia@university.edu', 'AIML', 1, 600, '2025-05-07 03:53:03', 45, 0),
(7, 'S007', 'Robert Miller', 'robert.miller@university.edu', 'AIDS', 1, 50, '2025-05-07 03:53:03', 5, 0),
(8, 'S008', 'Jennifer Davis', 'jennifer.davis@university.edu', 'EEE', 1, 150, '2025-05-07 03:53:03', 15, 0),
(9, 'S009', 'William Rodriguez', 'william.rodriguez@university.edu', 'ECE', 1, 250, '2025-05-07 03:53:03', 25, 0),
(10, 'S010', 'Elizabeth Martinez', 'elizabeth.martinez@university.edu', 'AERO', 1, 350, '2025-05-07 03:53:03', 30.5, 9),
(11, 'S011', 'Christopher Wilson', 'christopher.wilson@university.edu', 'MECH', 1, 750, '2025-05-13 04:37:09', 50, 10),
(12, 'S012', 'Amanda Anderson', 'amanda.anderson@university.edu', 'EI', 1, 820, '2025-05-13 04:37:09', 42, 10),
(13, 'S013', 'Matthew Taylor', 'matthew.taylor@university.edu', 'AUTO', 1, 900, '2025-05-13 04:37:09', 45, 0),
(14, 'S014', 'Ashley Thomas', 'ashley.thomas@university.edu', 'FOOD', 1, 650, '2025-05-13 04:37:09', 32.5, 0),
(15, 'S015', 'Daniel Hernandez', 'daniel.hernandez@university.edu', 'FASHION', 1, 3000, '2025-05-13 04:37:09', 66, 0);

-- --------------------------------------------------------

--
-- Table structure for table `student_subjects`
--

CREATE TABLE `student_subjects` (
  `id` int NOT NULL,
  `student_id` varchar(50) NOT NULL,
  `subject_code` varchar(10) NOT NULL,
  `semester` int NOT NULL,
  `year` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `subject_name` varchar(100) DEFAULT NULL,
  `department` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `student_subjects`
--

INSERT INTO `student_subjects` (`id`, `student_id`, `subject_code`, `semester`, `year`, `created_at`, `subject_name`, `department`) VALUES
(7, 'S001', 'SUB7', 2, 1, '2025-05-13 04:32:45', 'Mathematics2', 'IT'),
(8, 'S001', 'SUB8', 2, 1, '2025-05-13 04:32:45', 'Physics2', 'IT'),
(9, 'S001', 'SUB9', 2, 1, '2025-05-13 04:32:45', 'Chemistry2', 'IT'),
(10, 'S001', 'SUB10', 2, 1, '2025-05-13 04:32:45', 'Biology2', 'IT'),
(11, 'S001', 'SUB11', 2, 1, '2025-05-13 04:32:45', 'History2', 'IT'),
(12, 'S001', 'SUB12', 2, 1, '2025-05-13 04:32:45', 'Computer Science2', 'IT'),
(13, 'S003', 'SUB4', 1, 3, '2025-05-13 04:32:45', 'Biology', 'CSD'),
(14, 'S003', 'SUB2', 1, 3, '2025-05-13 04:32:45', 'Physics', 'CSD'),
(15, 'S003', 'SUB3', 1, 3, '2025-05-13 04:32:45', 'Chemistry', 'CSD'),
(16, 'S003', 'SUB5', 1, 3, '2025-05-13 04:32:45', 'History', 'CSD'),
(17, 'S003', 'SUB6', 1, 3, '2025-05-13 04:32:45', 'Computer Science', 'CSD'),
(18, 'S003', 'SUB7', 1, 3, '2025-05-13 04:32:45', 'Mathematics32', 'CSD'),
(19, 'S002', 'SUB8', 2, 2, '2025-05-13 04:32:45', 'Physics2', 'CSE'),
(20, 'S002', 'SUB9', 2, 2, '2025-05-13 04:32:45', 'Chemistry2', 'CSE'),
(21, 'S002', 'SUB10', 2, 2, '2025-05-13 04:32:45', 'Biology2', 'CSE'),
(22, 'S002', 'SUB11', 2, 2, '2025-05-13 04:32:45', 'History2', 'CSE'),
(23, 'S002', 'SUB12', 2, 2, '2025-05-13 04:32:45', 'Computer Science2', 'CSE'),
(24, 'S002', 'SUB14', 2, 2, '2025-05-13 04:32:45', 'Data Science Fundamentals', 'CSE'),
(39, 'S012', 'SUB5', 1, 4, '2025-05-13 04:37:18', 'History', 'EI'),
(40, 'S012', 'SUB7', 1, 4, '2025-05-13 04:37:18', 'Mathematics32', 'EI'),
(41, 'S012', 'SUB9', 1, 4, '2025-05-13 04:37:18', 'Chemistry2', 'EI'),
(42, 'S012', 'SUB11', 1, 4, '2025-05-13 04:37:18', 'History2', 'EI');

--
-- Triggers `student_subjects`
--
DELIMITER $$
CREATE TRIGGER `sync_department_insert` BEFORE INSERT ON `student_subjects` FOR EACH ROW BEGIN
    DECLARE student_dept VARCHAR(100);
    
    SELECT department INTO student_dept
    FROM student_rewards
    WHERE student_id = NEW.student_id;
    
    SET NEW.department = student_dept;
END
$$
DELIMITER ;
DELIMITER $$
CREATE TRIGGER `sync_department_update` BEFORE UPDATE ON `student_subjects` FOR EACH ROW BEGIN
    DECLARE student_dept VARCHAR(100);
    
    SELECT department INTO student_dept
    FROM student_rewards
    WHERE student_id = NEW.student_id;
    
    SET NEW.department = student_dept;
END
$$
DELIMITER ;
DELIMITER $$
CREATE TRIGGER `update_subject_name_on_insert` AFTER INSERT ON `student_subjects` FOR EACH ROW BEGIN
    UPDATE student_subjects ss
    JOIN subjects s ON ss.subject_code = s.code
    SET ss.subject_name = s.name
    WHERE ss.id = NEW.id;
END
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `student_subject_points`
--

CREATE TABLE `student_subject_points` (
  `id` int NOT NULL,
  `student_id` varchar(50) NOT NULL,
  `subject_code` varchar(10) NOT NULL,
  `subject_name` varchar(100) NOT NULL,
  `points` int NOT NULL DEFAULT '0',
  `marks` float NOT NULL DEFAULT '0',
  `section_points` json DEFAULT NULL,
  `semester` int NOT NULL DEFAULT '1',
  `year` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `student_subject_points`
--

INSERT INTO `student_subject_points` (`id`, `student_id`, `subject_code`, `subject_name`, `points`, `marks`, `section_points`, `semester`, `year`, `created_at`, `updated_at`) VALUES
(59, 's001', 'SUB7', 'Mathematics2', 70, 4, '[67, 0, 0, 0]', 1, 2025, '2025-05-20 11:08:21', '2025-05-23 09:57:25'),
(60, 's001', 'SUB10', 'Biology2', 60, 5.5, '[50, 10, 0, 0]', 1, 2025, '2025-05-20 11:08:29', '2025-05-20 11:08:29'),
(61, 's003', 'SUB4', 'Biology', 240, 10.5, '[50, 60, 80, 30]', 1, 2025, '2025-05-22 04:10:13', '2025-05-22 04:10:13'),
(62, 's003', 'SUB2', 'Physics', 250, 11, '[50, 60, 80, 60]', 1, 2025, '2025-05-22 04:10:35', '2025-05-22 04:11:17'),
(64, 's002', 'SUB8', 'Physics2', 235, 9, '[75, 90, 60, 0]', 1, 2025, '2025-05-23 04:51:03', '2025-05-23 04:51:03'),
(65, 's002', 'SUB12', 'Computer Science2', 205, 8.5, '[75, 90, 30, 0]', 1, 2025, '2025-05-23 04:51:13', '2025-05-23 05:08:31'),
(66, 's002', 'SUB14', 'Data Science Fundamentals', 185, 8, '[75, 90, 0, 0]', 1, 2025, '2025-05-23 04:51:17', '2025-05-23 04:51:25'),
(68, 's002', 'SUB9', 'Chemistry2', 245, 9, '[75, 90, 60, 0]', 1, 2025, '2025-05-23 05:08:09', '2025-05-23 05:08:09'),
(71, 's001', 'SUB11', 'History2', 10, 0, '[7, 0, 0, 0]', 1, 2025, '2025-05-23 09:57:36', '2025-05-23 09:57:36'),
(72, 's002', 'SUB8', 'Physics2', 250, 9, '[75, 90, 60, 0]', 1, 1, '2025-05-24 05:03:09', '2025-05-24 05:03:09'),
(73, 's002', 'SUB11', 'History2', 250, 9, '[75, 90, 60, 0]', 1, 1, '2025-05-24 05:03:25', '2025-05-24 05:03:25'),
(74, 's001', 'SUB7', 'Mathematics2', 23, 1, '[22, 0, 0, 0]', 1, 1, '2025-05-24 06:30:33', '2025-05-24 07:11:41');

-- --------------------------------------------------------

--
-- Table structure for table `student_totals`
--

CREATE TABLE `student_totals` (
  `student_id` varchar(50) NOT NULL,
  `total_points` int NOT NULL DEFAULT '0',
  `total_marks` float NOT NULL DEFAULT '0',
  `points_remaining` int NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `student_totals`
--

INSERT INTO `student_totals` (`student_id`, `total_points`, `total_marks`, `points_remaining`) VALUES
('s001', 0, 0, 0),
('s004', 1900, 10.5, 0),
('s010', 0, 0, 0);

-- --------------------------------------------------------

--
-- Table structure for table `subjects`
--

CREATE TABLE `subjects` (
  `id` int NOT NULL,
  `code` varchar(10) NOT NULL,
  `name` varchar(100) NOT NULL,
  `semester` int NOT NULL DEFAULT '1',
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `subject_type` enum('core','addon','elective') NOT NULL DEFAULT 'core'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `subjects`
--

INSERT INTO `subjects` (`id`, `code`, `name`, `semester`, `is_active`, `created_at`, `subject_type`) VALUES
(1, 'SUB1', 'Mathematics', 1, 1, '2025-05-07 05:51:01', 'core'),
(2, 'SUB2', 'Physics', 1, 1, '2025-05-07 05:51:01', 'core'),
(3, 'SUB3', 'Chemistry', 1, 1, '2025-05-07 05:51:01', 'core'),
(4, 'SUB4', 'Biology', 1, 1, '2025-05-07 05:51:01', 'core'),
(5, 'SUB5', 'History', 1, 1, '2025-05-07 05:51:01', 'elective'),
(6, 'SUB6', 'Computer Science', 1, 1, '2025-05-07 05:51:01', 'addon'),
(7, 'SUB7', 'Mathematics2', 1, 1, '2025-05-07 05:51:01', 'core'),
(8, 'SUB8', 'Physics2', 1, 1, '2025-05-07 05:51:01', 'core'),
(9, 'SUB9', 'Chemistry2', 1, 1, '2025-05-07 05:51:01', 'core'),
(10, 'SUB10', 'Biology2', 1, 1, '2025-05-07 05:51:01', 'core'),
(11, 'SUB11', 'History2', 1, 1, '2025-05-07 05:51:01', 'elective'),
(12, 'SUB12', 'Computer Science2', 1, 1, '2025-05-07 05:51:01', 'addon'),
(13, 'SUB13', 'Advanced Programming', 1, 1, '2025-05-13 05:37:01', 'addon'),
(14, 'SUB14', 'Data Science Fundamentals', 1, 1, '2025-05-13 05:37:01', 'addon'),
(15, 'SUB15', 'Digital Marketing', 1, 1, '2025-05-13 05:37:01', 'elective'),
(16, 'SUB16', 'Environmental Science', 1, 1, '2025-05-13 05:37:01', 'elective');

-- --------------------------------------------------------

--
-- Table structure for table `subject_points`
--

CREATE TABLE `subject_points` (
  `student_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `subject_code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `semester` int NOT NULL DEFAULT '1',
  `year` int NOT NULL DEFAULT '1',
  `subject_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `points` int NOT NULL DEFAULT '0',
  `marks` float NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `student_email` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `subject_points`
--

INSERT INTO `subject_points` (`student_id`, `subject_code`, `semester`, `year`, `subject_name`, `points`, `marks`, `created_at`, `updated_at`, `student_email`) VALUES
('s002', 'SUB1', 1, 1, 'Mathematics', 45, 4.5, '2025-05-14 09:41:10', '2025-05-14 09:41:10', ''),
('s003', 'SUB4', 1, 1, 'Biology', 240, 10.5, '2025-05-14 09:22:49', '2025-05-14 09:22:49', ''),
('s007', 'SUB1', 1, 1, 'Mathematics', 45, 4.5, '2025-05-14 09:24:19', '2025-05-14 09:24:19', ''),
('s008', 'SUB3', 1, 1, 'Chemistry', 40, 4, '2025-05-14 07:15:33', '2025-05-14 07:15:33', '');

-- --------------------------------------------------------

--
-- Table structure for table `yearly_config`
--

CREATE TABLE `yearly_config` (
  `year` int NOT NULL,
  `config_json` text NOT NULL,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `config`
--
ALTER TABLE `config`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `config_key` (`config_key`);

--
-- Indexes for table `config_year`
--
ALTER TABLE `config_year`
  ADD PRIMARY KEY (`year`,`department`);

--
-- Indexes for table `departments`
--
ALTER TABLE `departments`
  ADD PRIMARY KEY (`code`);

--
-- Indexes for table `redemption_dates`
--
ALTER TABLE `redemption_dates`
  ADD PRIMARY KEY (`year`,`department`);

--
-- Indexes for table `reward_config`
--
ALTER TABLE `reward_config`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `config_key` (`config_key`);

--
-- Indexes for table `student_rewards`
--
ALTER TABLE `student_rewards`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `student_id` (`student_id`),
  ADD UNIQUE KEY `student_id_2` (`student_id`),
  ADD UNIQUE KEY `student_id_3` (`student_id`),
  ADD UNIQUE KEY `uk_student_department` (`student_id`,`department`),
  ADD KEY `idx_student_id` (`student_id`);

--
-- Indexes for table `student_subjects`
--
ALTER TABLE `student_subjects`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `student_id` (`student_id`,`subject_code`,`semester`,`year`),
  ADD KEY `subject_code` (`subject_code`),
  ADD KEY `fk_student_department` (`student_id`,`department`);

--
-- Indexes for table `student_subject_points`
--
ALTER TABLE `student_subject_points`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `student_id` (`student_id`,`subject_code`,`semester`,`year`),
  ADD KEY `subject_code` (`subject_code`);

--
-- Indexes for table `student_totals`
--
ALTER TABLE `student_totals`
  ADD PRIMARY KEY (`student_id`);

--
-- Indexes for table `subjects`
--
ALTER TABLE `subjects`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`);

--
-- Indexes for table `subject_points`
--
ALTER TABLE `subject_points`
  ADD PRIMARY KEY (`student_id`,`subject_code`);

--
-- Indexes for table `yearly_config`
--
ALTER TABLE `yearly_config`
  ADD PRIMARY KEY (`year`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `config`
--
ALTER TABLE `config`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=697;

--
-- AUTO_INCREMENT for table `reward_config`
--
ALTER TABLE `reward_config`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=344;

--
-- AUTO_INCREMENT for table `student_rewards`
--
ALTER TABLE `student_rewards`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- AUTO_INCREMENT for table `student_subjects`
--
ALTER TABLE `student_subjects`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=59;

--
-- AUTO_INCREMENT for table `student_subject_points`
--
ALTER TABLE `student_subject_points`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=76;

--
-- AUTO_INCREMENT for table `subjects`
--
ALTER TABLE `subjects`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `student_subjects`
--
ALTER TABLE `student_subjects`
  ADD CONSTRAINT `fk_student_department` FOREIGN KEY (`student_id`,`department`) REFERENCES `student_rewards` (`student_id`, `department`) ON UPDATE CASCADE,
  ADD CONSTRAINT `student_subjects_ibfk_1` FOREIGN KEY (`student_id`) REFERENCES `student_rewards` (`student_id`),
  ADD CONSTRAINT `student_subjects_ibfk_2` FOREIGN KEY (`subject_code`) REFERENCES `subjects` (`code`);

--
-- Constraints for table `student_subject_points`
--
ALTER TABLE `student_subject_points`
  ADD CONSTRAINT `student_subject_points_ibfk_1` FOREIGN KEY (`student_id`) REFERENCES `student_rewards` (`student_id`),
  ADD CONSTRAINT `student_subject_points_ibfk_2` FOREIGN KEY (`subject_code`) REFERENCES `subjects` (`code`);

--
-- Constraints for table `student_totals`
--
ALTER TABLE `student_totals`
  ADD CONSTRAINT `student_totals_ibfk_1` FOREIGN KEY (`student_id`) REFERENCES `student_rewards` (`student_id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
