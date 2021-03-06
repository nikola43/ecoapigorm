-- phpMyAdmin SQL Dump
-- version 4.6.6deb5ubuntu0.5
-- https://www.phpmyadmin.net/
--
-- Servidor: localhost:3306
-- Tiempo de generación: 24-02-2021 a las 22:06:31
-- Versión del servidor: 5.7.33-0ubuntu0.18.04.1
-- Versión de PHP: 7.2.24-0ubuntu0.18.04.7

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `ecopanel`
--
CREATE DATABASE IF NOT EXISTS `ecopanel` DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;
USE `ecopanel`;

DELIMITER $$
--
-- Funciones
--
CREATE DEFINER=`root`@`localhost` FUNCTION `F_COUNT_USERS_BY_CLINIC` (`clinic_id` INTEGER) RETURNS INT(11) BEGIN
  DECLARE usersByClinic INT unsigned DEFAULT 0;

  SELECT COUNT(u.id)
  INTO usersByClinic
        FROM clinics as c
        INNER JOIN users as u
        ON c.id = u.clinic_id
        AND c.available = 1
        AND u.available = 1
        AND c.id = clinic_id;
  RETURN usersByClinic;
END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `calculators`
--

CREATE TABLE `calculators` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `week` int(11) NOT NULL DEFAULT '12',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `calculators`
--

INSERT INTO `calculators` (`id`, `user_id`, `week`, `updated_at`, `created_at`) VALUES
(1, 1, 28, '2020-11-25 13:12:08', '2020-11-23 00:00:00'),
(52, 2, 28, '2020-11-25 13:12:08', '2020-11-23 00:00:00'),
(53, 4, 25, '2021-02-24 21:26:55', '2020-11-23 00:00:00');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `calculator_details`
--

CREATE TABLE `calculator_details` (
  `id` int(11) NOT NULL,
  `image_url` varchar(256) NOT NULL,
  `text` varchar(4096) CHARACTER SET utf8 COLLATE utf8_spanish_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `calculator_details`
--

INSERT INTO `calculator_details` (`id`, `image_url`, `text`) VALUES
(1, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/2SM.jpg', 'Convenciona-se que a contagem da idade gestacional começa mesmo antes de ter acontecido a concepção (fecundação). Dessa forma, a semana 1 começa com o 1º dia da sua última menstruação – mesmo ainda não estando grávida. Para as mulheres com ciclos de 28 dias, a ovulação acontece por volta do 14º dia do ciclo. Esse é um bom período para engravidar. '),
(2, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/2SM.jpg', 'O óvulo já está quase pronto para abandonar o ovário e seguir pela trompa de Falópio ao encontro dos espermatozoides: o tão esperado momento da fecundação. Entre os dias 12 e 16 do ciclo (para mulheres com ciclo regular de 28 dias) estará no seu período fértil.'),
(3, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/3SM.jpg', 'Esta semana é fulcral dado que vai ocorrer a fecundação do óvulo, marcando o início de uma das viagens mais especiais para a mamã e papá.\nO óvulo fecundado divide-se pela primeira vez passado apenas um dia da fecundação. A divisão continua nos dias seguintes, enquanto viaja pela trompa de Falópio, até se formar uma estrutura que se designa blastocisto que nos primeiros quatro dias pode ter mais de 100 células. Ao 5º dia as células começam a separar-se em dois grupos – um para formar a placenta e outro para formar o bebé. As células que vão formar o bebé são as chamadas células estaminais que têm a capacidade extraordinária de se transformarem em mais de 200 classes células de qualquer parte do corpo (por isso são tão interessantes para investigar tratamentos de determinadas doenças).\nOs 23 cromossomas do óvulo encontram-se unidos aos 23 cromossomas do espermatozoide. Com esta união é decidido o sexo do bebé, características físicas e mentais do bebé.\nAo 7º dia o blastocisto é capaz de se aderir ao endométrio – esse processo é denominado nidação. \nAté notar o atraso na menstruação, a mamã não saberá que já está grávida. Também existem mamãs que notam já nesta fase alterações hormonais. Seja qual for o seu caso, está grávida. Parabéns!\n\nGémeos: Se dois óvulos são libertados pelos ovários e são fecundados por dois espermatozóides diferentes, serão gerados gémeos fraternos (dizigóticos). Se um óvulo é fecundado e depois se divide em dois, serão gerados os gémeos idênticos.\n'),
(4, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/4SM.jpg', 'Nesta semana deveria vir a menstruação, mas se engravidou ela não aparecerá. Algumas mulheres podem ter um pequeno sangramento que ocorre quando o óvulo fecundado se fixa na parede do útero. Mas o volume desse sangramento na nidação é bem menor que uma menstruação.\nAlguns dos primeiros sintomas de gravidez são muito semelhantes aos que costuma sentir no período pré-menstrual: cansaço, sensibilidade mamária e alterações de humor. Ao sentir esses sintomas pode até ficar desanimada por achar que ainda não foi desta vez. Portanto, para ter certeza espere mais alguns dias…\nAo fim da quarta semana de gestação o seu bebé mede de 0,36 a 1 mm, da cabeça às nádegas, mas ainda não pode ser visto numa ecografia. Talvez seja possível ver apenas uma bolinha preta dentro do útero. Essa bolinha é o saco gestacional.\n'),
(5, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/5SM.jpg', 'Chegou o tão esperado positivo no teste de gravidez!\nDurante estas primeiras e cruciais semanas, os órgãos vitais, ossos e sistema nervoso iniciam o seu desenvolvimento. \nO embrião dividiu-se em três camadas, que darão lugar ao futuros órgãos do bebé: a ectoderme, que vai dar lugar ao sistema nervoso central; a endoderme, que vai dar lugar aos intestinos, pâncreas, fígado e glândula tiroide; e a mesoderme, de onde surgirão os ossos, músculos e sistema sanguíneo. \nO embrião vai esboçar logo a cabeça, os olhos, as orelhas e restantes membros.\nO tubo neuronal, que se irá transformar no cérebro e na medula espinal, também se começa a formar. \nPor outro lado, o tecido embrionário começou a formar a estrutura que vai acabar por ser o coração do bebé.\nA placenta que é o órgão que vai nutrir o bebé nos próximos meses, também se está a formar. \nO volume de sangue da mamã aumenta 50%, para corresponder à necessidade de oxigénio do bebé. No interior do seu corpo, está a criar um mundo seguro para que o seu pequenino possa crescer e desenvolver, protegido do mundo exterior.\nNesta semana de gravidez, o embrião mede entre 1 e 2 mm.\n'),
(6, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/6SM.jpg', 'O coração do bebé é um grupo de células musculares inicialmente adormecidas, mas por volta dos 21 dias explode de vida. Uma célula contrai-se espontaneamente e contagia as células vizinhas começando uma reacção em cadeia até que todas as células do coração começam a bater. Não é fantástico? Estas células estão programadas para controlar os batimentos do coração do bebé até ao cérebro estar suficientemente desenvolvido para o poder fazer. O coração vai distribuir o alimento e o oxigénio que o bebé precisa para crescer.\nO coração bate, em média, 150 vezes por minuto. Quase o dobro de um adulto!\nAlém disso, o vínculo com o seu bebé torna-se mais intenso, pois os vasos sanguíneos que unem o cordão umbilical com a placenta, aparecem.\nO tubo neural – que será o cérebro e a medula espinal – está a começar a fechar-se.\nNa sua cabeça vão aparecer pequenos buracos lado a lado que, depois, serão os seus olhos; a sua cabeça vai inclinar-se para o corpo e na sua cabeça e pescoço vão esboçar-se a mandíbula inferior e a sua laringe.\nE ainda há mais! Na parte inferior vai desenvolver-se uma protuberância parecida com uma cauda que, depois, dará lugar às suas pernas. Na parte superior também vão aparecer os primórdios dos seus braços.\nO coração e o fígado combinados possuem o mesmo volume da cabeça nesse período.\nEle mede cerca de 4 mm de comprimento (da cabeça às nádegas) nesta semana e tem a forma de um “C”. O crescimento é muito rápido nessa fase.\n'),
(7, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/7SM.jpg', 'Durante estas próximas semanas o embrião vai assumir uma forma cada vez mais parecida com a de um bebé.\nNa sétima semana de gestação o embrião cresce muito, pois no início da semana mede entre 4 e 5 mm do cimo da cabeça às nádegas e no final da mesma duplicou o seu comprimento chegando aos 8 ou 11 mm.\nO cérebro divide-se em dois hemisférios e vai aumentando o seu tamanho, o coração dividiu-se em duas câmaras, os pulmões têm um brônquio primário para permitir a passagem do ar, e começam a aparecer as fossas nasais e as órbitas dos olhos.\nA placenta continua o seu processo de estabilização, embora ainda não esteja preparada para servir de fonte de alimentação.\nDurante esta semana, os órgãos vitais do embrião, como os pulmões e o intestino, começam a desenvolver-se, o fígado começa a produzir glóbulos brancos e o pâncreas a segregar insulina.\nNo final desta semana o embrião apresenta o tubérculo genital, que originará os órgãos genitais externos (pénis e bolsa escrotal no homem; clitóris, grandes lábios e parte da vagina, na mulher), mas ainda é cedo para se perceber numa ecografia.\n'),
(8, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/8SM.jpg', 'O bebé mede cerca de 8-11 mm de comprimento (da cabeça às nádegas) no início desta semana. \nA sua cabeça, onde já se começam a formar as orelhas, o ouvido interno e os olhos é ainda muito grande em comparação com o resto do corpo. Por isso, embora a sua coluna se vá endireitando pouco a pouco, a cabeça ainda permanece inclinada para a frente.\nNesta fase, o seu cérebro é uma estrutura unida à medula espinal, mas que agora se começa a torcer e a formar áreas distintas.\nNesta semana também vão formar-se o lábio inferior e a mandibular. O bebé já apresenta narinas. \nPor outro lado, a cauda do embrião começa a desaparecer e vai acabar por ser substituída pelos ossos que vão formar a parte inferior da coluna do bebé.\nOs dedos dos seus pés e mãos começam a formar-se, embora ainda permaneçam unidos e os cotovelos já são visíveis.\nA maioria dos órgãos do bebé – rins, fígado, coração, cérebro – já está formada. As células ósseas começam a substituir as cartilagens e começam a ser formadas as articulações. \nPor incrível que pareça, os intestinos começam-se a desenvolver dentro do cordão umbilical e migrarão depois para o abdómen do seu bebé quando ele estiver bastante grande para acomodá-los. \nDurante esta semana, uma ecografia poderá confirmar a idade gestacional do bebé e detetar o batimento do coração, onde já se podem diferenciar as válvulas aórtica e pulmonar e os cerca de 150 batimentos por minuto.\nA placenta assume agora a função de alimentar o bebé através do cordão umbilical. Irá receber nutrientes, oxigénio e água.\nNo final desta semana o bebé mede cerca de 13-17 mm (da cabeça às nádegas) e pesa em torno de 1 g.\n'),
(9, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/9SM.jpg', 'As suas mãos e dedinhos já começam a aparecer e as suas pernas crescem, formando-se também os pés na extremidade. Além disso, o seu rosto vai começar a adquirir um aspeto mais parecido com o nosso e, embora ainda não o note, numa ecografia já se pode identificar os primeiros movimentos do seu pequenino.\nOs olhos, que até agora estavam na parte lateral da cabeça, começam a deslocar-se para a parte frontal. As pálpebras vão unir-se e permanecerão fechadas até por volta da 26º semana de gestação.\nAs aberturas nasais e a ponta do nariz estão completamente formadas. \nOs ossos e cartilagens continuam a desenvolverem-se. \nO coração está completamente formado; o diafragma separa o tórax do abdómen. \nAlgumas glândulas já começam a funcionar e são detectadas as primeiras ondas cerebrais. \nNesta semana todos os órgãos principais do sistema digestivo do bebé começaram a desenvolver-se, embora ainda não possam realizar nenhuma função digestiva.\nNo final desta semana, o embrião mede cerca de 23 mm (da cabeça às nádegas) e pesa cerca de 2 g.\n'),
(10, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/10SM.jpg', 'Nesta semana, embora o seu bebé ainda pareça um peixinho de aproximadamente 31 mm de comprimento (da cabeça às nádegas) e 4 g de peso (peso de dois clipes de papel), os seus órgãos genitais começam a formar-se, as articulações principais dos ombros, dos cotovelos, da bacia e dos joelhos já são visíveis e os órgãos internos já estão nos seus devidos lugares, embora ainda lhes falte bastante tempo para funcionarem com autonomia. \nAlém disso, a língua, a laringe e a tiróide começam a sua formação.\nNesta semana, a cabeça do bebé separa-se ligeiramente do peito, que contribui para o progressivo crescimento do pescoço e do maxilar.\nO seu nariz já sobressai na cara, a boca e os lábios desenvolveram-se quase completamente, tal como o ouvido externo, que já está completamente formado, embora ainda não ocupe a sua posição final.\nOutras mudanças substanciais no bebé durante esta décima semana de gravidez são o crescimento dos dedos das mãos e dos pés, bem como a sua separação, a formação dos pulsos e o facto de se poderem dobrar, e o início da formação da língua, do palato e dos primórdios dos seus dentinhos.\nNesta última fase de desenvolvimento embrionário, todas as estruturas externas e internas essenciais estão presentes. Os principais sistemas estão integrados e formados. \n'),
(11, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/11SM.jpg', 'O bebé adquiriu uma aparência completamente humana, apesar da sua cabeça representar ainda metade de seu tamanho.\nO bebé tem cerca de 41 mm de comprimento (da cabeça às nádegas) e pesa aproximadamente 7 g.\nDurante esta semana o bebé começa a gerar os seus próprios glóbulos vermelhos. \nEm breve vai começar a produzir urina, o principal componente do líquido amniótico, onde o bebé vai começar a estar muito ativo, dando pontapés e mexendo-se, embora a mamã ainda não o sinta.\nNa boca do bebé já se formou um palato duro.\nOs ovários ou testículos começam a formar-se e, no caso dos meninos, os seus testículos vão começar a produzir testosterona, a hormona masculina.\nA bexiga e o reto do bebé já se separaram e o diafragma, já completo, permite ao bebé realiar movimentos respiratórios.\n'),
(12, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/12SM.jpg', 'A 12ª semana de gravidez marca o final do primeiro trimestre. O bebé foi promovido a feto! Já completou a primeira fase do seu desenvolvimento!\n \nCom esta semana termina o período de “embriogénese”, durante o qual o embrião está mais exposto aos perigos que poderiam prejudicar a correta formação dos diferentes órgãos. \nO cérebro continua o seu desenvolvimento.\nNesta semana o seu coração bate muito rapidamente, a umas 160 pulsações por minuto, o que é o dobro do da mãe.\nOs seus membros estão completamente formados e cresceram, a sua cabeça tem uma forma muito mais arredondada, a sua boca já se pode abrir e fechar, as suas orelhas já estão praticamente no local correcto e os seus intestinos, que até agora estavam presos à base do cordão umbilical, deslocam-se para a cavidade abdominal, onde já há espaço para eles. \nOs dedos das mãos e pés já se separaram e os pelos e unhas iniciam o seu crescimento. \nO aparelho genital externo do bebé começa a definir-se, mas ainda é cedo para se ver com exactidão através da ecografia. \nO líquido amniótico começa a acumular-se à medida que os rins do bebé começam a produzir urina. \nOs músculos das paredes intestinais começam-se a movimentar que ajudam na digestão e movimentação dos alimentos. \nO bebé mede cerca de 54 mm (da cabeça às nádegas) e pesa em torno de 14 g. \nNa 12ª semana de gravidez a maioria das mulheres realiza a primeira ecografia das três principais que se fazem numa gravidez saudável. Nesta ecografia será determinada a espessura da translucência da nuca, uma zona que fica atrás do pescoço que se considera um marcador de alterações genéticas como a síndrome de Down ou a de Turner.\n'),
(13, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/13SM.jpg', 'Acabou de entrar oficialmente no 2º trimestre!\nOs incómodos típicos do início da gestação aliviam e o bebé prossegue com o seu rápido desenvolvimento, especialmente no que diz respeito ao cérebro e ao controlo muscular.\nTodas as suas articulações já estão formadas, o que lhe permite um amplo leque de movimentos.\nO bebé mede cerca de 74 mm (da cabeça às nádegas) e pesa aproximadamente 23 g.\nOs braços e as pernas alongam-se rapidamente, as articulações das ancas maturam e os dedos dos pés separam-se.\nO seu corpo não tem gordura e os seus ossos começam a notar-se sob a sua delicada pele.\nNeste momento, o bebé está muito ativo na bolsa amniótica, onde flutua e tem espaço suficiente para se mexer com liberdade.\nAgora a testa do bebé está elevada. Nela podem ver-se as uniões das placas de osso que compõem o crânio. Os hemisférios cerebrais esquerdo e direito começam a ligar-se e as primeiras a maturar serão as fibras motoras, as que controlam os movimentos. Depois serão os nervos sensoriais, que são os que controlam a alimentação. O desenvolvimento do cérebro acelerar-se-á significativamente durante as próximas três semanas e ficará completo em cerca de dez.\nÉ possível agora, em alguns casos, determinar o sexo do bebé, olhando para o tubérculo genital e observando sua angulação. \n'),
(14, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/14SM.jpg', 'Na 14ª semana de gravidez, o bebé já adquiriu uma aparência completamente humana. Os seus traços faciais já estão completamente definidos, os olhos e as orelhas já estão na sua localização final e aparecem as sobrancelhas e os seus primeiros cabelos.\nNeste momento, o bebé recebe todo o seu alimento a partir da placenta, que já tem um tamanho maior do que ele.\nAgora, o bebé mede cerca de 87 mm de comprimento (da cabeça às nádegas) e pesa em torno de 43 g.\nO cérebro continua a desenvolver-se rapidamente.\nO sistema nervoso central do bebé, que inclui o cérebro e a coluna, já tem os seus componentes básicos e as ligações entre células nervosas individuais tornam-se mais organizadas.\nNa 14ª semana de gravidez, a bexiga do bebé enche-se e esvazia-se a cada 30 minutos. Primeiro engole o líquido amniótico, filtra-o através dos rins e depois liberta-o em forma de urina.\nO bebé produz agora uma pequena quantidade de glóbulos brancos, embora ainda dependa da mamã para combater as infeções e os seus glóbulos vermelhos contêm hemoglobina, que serve para transportar o oxigénio às células do corpo. O seu sistema sanguíneo já é capaz de criar e dissolver coágulos sanguíneos.\nComeça a fazer movimentos respiratórios – inspiração e expiração. \nO pescoço está-se alongando e o queixo ainda repousa sobre o tórax. \nAs mãos estão-se a tornar funcionais e o bebé começa a aprender a movê-las (mais como um reflexo do que por vontade). \nÉ a partir desta fase que a determinação do sexo através de uma ecografia já tem maior exactidão. \n'),
(15, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/15SM.jpg', 'Nesta fase, o bebé mede cerca de 10 cm do cimo da cebeça até às nádegas e pesa em torno de 70 g. Durante as próximas semanas o seu tamanho vai duplicar e, agora, o crescimento dos ossos e dos músculos avança a um ritmo muito rápido. \nA pele é muito fina e transparente; é possível verem-se vasos sanguíneos através da mesma.\nO bebé cobre-se de um pelo muito fino denominado lanugo que continuará a crescer até à 26ª semana de gestação.\nA medula espinal já está totalmente formada e estende-se por todo o canal vertebral, com nervos que saem em cada vértebra.\nO seu pescoço alonga-se e já pode levantar mais a cabeça e separar a cabeça do peito.\nOs antebraços, pulsos, mãos e dedos já se diferenciaram claramente.\nOs seus braços cresceram o suficiente para aproximar as mãozinhas da cara e, além disso, é possível vê-lo na ecografia a chuchar no dedo.\nAs feições individualizam-se e as expressões faciais começam a ficar evidentes.\n'),
(16, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/16SM.jpg', 'Numa gravidez de 16 semanas o bebé já tem uma grande mobilidade nas suas mãos e é capaz de as abrir e fechar.\nNesta semana o bebé mede cerca de 11,6 cm (da cabeça às nádegas) e pesa em média 100 g.\nJá poderá ser possível sentir vibrações relacionadas com as cambalhotas e pontapés que o bebé vai dando. As pernas são maiores que os braços e ele move os membros com muita frequência.\nSe é o seu primeiro filho, talvez ainda não se aperceba dos movimentos. É perfeitamente normal.\nO bebé também começa a desenvolver folículos pilosos. As unhas continuam a crescer.\nNo seu cérebro formam-se as células nervosas, embora a sua atividade neuronal ainda seja muito imatura, tal como acontece com o intestino, qua apesar de poder dar conta, por vezes, de pequenas quantidades de líquido amniótico que o bebé engole, ainda é demasiado imaturo para funcionar adequadamente de forma regular.\nJá é possível perceber a existência de mecónio nos intestinos. \nA vagina e o ânus já estão abertos\n'),
(17, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/17SM.jpg', 'Na 17ª semana de gravidez o bebé tem muito espaço para se mexer e não para quieto. Estica-se, retorce-se e muda constantemente de posição. \nO minúsculo coração está bombeando cerca de 21 litros de sangue por dia! \nOs reflexos estão funcionantes e o bebé já é capaz também de sugar e deglutir. Periodicamente engole o líquido amniótico onde está mergulhado e elimina-o através da urina.\nNa 17ª semana de gravidez o bebé mede cerca de mede cerca de 13 cm de comprimento (da cabeça às nádegas)  e pesa cerca de 140 g.\nOs seus pulmões continuam a ramificar-se: as células que cobrem as vias respiratórias produzem de maneira constante um fluido que abandona os pulmões quando o bebé efetua movimentos respiratórios, um movimento regulado pelas cordas vocais situadas na laringe.\nDentro da boca, as papilas gustativas amadureceram, embora o bebé não possa ainda notar qualquer sabor pois as ligações nervosas ainda são imaturas.\nO bebé já pode ouvir alguns sons de fora do corpo da mãe, embora a audição ainda não esteja completamente desenvolvida.\n'),
(18, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/18SM.jpg', 'A las 18 semanas de embarazo, el bebé puede bostezar y hacer gestos faciales. Incluso podrá tener hipo. Ya le funcionan las cuerdas vocales y podría llorar. Es posible que por fin empieces a sentir algunos movimientos del bebé porque empieza a patear y a mover sus manos con más fuerza. Si ya has tenido un bebé, reconocerás esos movimientos enseguida. Si es tu primer embarazo, quizás no los notes hasta dentro de dos semanas. Sus ojos y orejas ya alcanzado su ubicación definitiva en esta semana.\r\nAdemás, los huesos del oído interno y las terminaciones nerviosas del cerebro se han desarrollado lo suficiente. El bebé puede empezar a escuchar sonidos como sus latidos o el trayecto de la sangre a través del cordón umbilical.\r\nCerca de esta semana vendría de la segunda ecografía. En ella podremos analizar la anatomía del bebé y su ritmo de crecimiento. Te servirá para predecir posibles complicaciones.\r\nSus medidas: el bebé seguirá creciendo esta semana, aunque a un ritmo más lento. Tendrá unos 14 centímetros y 150 gramos.'),
(19, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/19SM.jpg', 'En la semana 19 de embarazo el sistema nervioso, en especial el cerebro, se está desarrollando y creciendo rápidamente.\r\nPor otro lado, el bebé ya hace movimientos activos que son más fáciles de sentir. Te darás cuenta de que hay momentos en los que el bebé parece estar dormido y otros en que se mueve mucho. Esto se debe a que duerme más y se despierta con más energía.\r\n\r\n\r\n\r\n\r\nConocer el sexo del bebé es más fácil y con resultados más fiables a partir de esta semana. Si es una niña sus ovarios ya contienen huevos primitivos. Las orejas se acercan a su posición final.\r\nSus medidas: el bebé mide unos 15 centímetros.'),
(20, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/20SM.jpg', '¡Ya has llegado a la mitad del embarazo! Se suele hablar de 40 semanas de embarazo a partir de la última menstruación. Sin embargo, la duración puede variar de 37 a 42 semanas. Durante la segunda mitad del embarazo, el peso del bebé aumentará más de diez veces (unos 3 kilos) y su longitud será el doble que en esta semana (pasará de unos 25 centímetros a unos 50).\r\nEl bebé puede oír a partir de las 2 o semanas de embarazo. El líquido amniótico que lo rodea distorsiona los sonidos (como cuando estamos debajo del agua), pero aún así el bebé ya te puede escuchar. Puede reconocer una música, el latir de su corazón o tu respiración. Además, percibe la luz, se mueve, traga, orina, quizás comienza a tener memoria, etc.\r\nLos pulmones y el tubo digestivo del bebé están madurando. Su cerebro cuenta con 30.000 millones de neuronas y está desarrollando especialmente áreas destinadas a los sentidos del gusto, el olfato, la audición, la visión y el tacto. Si es una niña, sus ovarios ya cuentan con 6 veces más óvulos que al momento de nacer. Un millón de óvulos aproximadamente, tiene una niña al momento de nacer.\r\nTu bebé ahora tiene cejas delgadas, pelo en la cabeza y miembros muy bien desarrollados. La forma y las proporciones generales del bebé son completamente humanas. Los movimientos son fundamentales para que no haya deformidades articulares ni corporales.\r\nCualquier pérdida durante estas primeras 20 semanas se denominará aborto, ya que el nacimiento en esta edad fetal es uniformemente letal, debido a la ausencia estructural de la porción respiratoria de los pulmones.\r\nSus medidas: el bebé mide aproximadamente 20 cm largo y pesa cerca de 255 gramos, ¡aún le queda mucho por crecer!'),
(21, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/21SM.jpg', 'En estos días puede que hayas notado que tu bebé no para de moverse. A partir de la semana 21 de embarazo, los movimientos del pequeño serán muy frecuentes. Si tu bebé es una niña, ya tiene la vagina formada, y si está en la posición adecuada, te podrán confirmar si será un precioso varón o una dulce hembra cuando te hagan una ecografía (ultrasonido), en caso de que todavía no lo sepas.\r\nTodos esos movimientos sirven para estimular su desarrollo físico y mental. Quizás durante el día no sientas las pataditas, giros, estiramientos y toda la gimnasia que tu bebe hace, pero por la noche… no te va a dejar dormir.\r\nPero, ¿por qué espera a que estés descansando para moverse tanto? Lo cierto es que durante el día también se mueve igual, pero tú sientes menos toda esa actividad que cuando estás descansando sin moverte.\r\nSus medidas: tu bebé ya tiene la longitud de una zanahoria. Ahora mide casi 27 centímetros desde la cabecita hasta los pies y además, ¡ya pesa unos 330 gramos!'),
(22, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/22SM.jpg', 'La piel delgada y rojiza del bebé está cubierta por una sustancia gruesa, blanca y cerosa llamada vermix o unto sebáceo. Ésta protege la piel del bebé contra las sustancias del líquido amniótico.\nEl sistema límbico del bebé está en pleno desarrollo y será el encargado de controlar los sentimientos y las emociones. Aunque parezca increíble, ésto le permitirá tener cambios en el estado de ánimo en las próximas semanas. Las uñas, párpados y cejas del bebé son visibles. Sólo faltan las pestañas.\nSus medidas: el bebé mide 20 centímetros y pesa 340 gramos.\n'),
(23, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/23SM.jpg', 'En la semana 23 de embarazo se está depositando el pigmento que colorea su piel. Ésta tiene una apariencia arrugada que se alisará en las próximas semanas. Al mismo tiempo, empezará a desarrollar su cerebro con rapidez. Sus medidas empiezan a ser más proporcionadas.\r\nTodos los sistemas del bebé (digestivo, circulatorio y respiratorio) están madurando y preparándose para la vida fuera del útero.\r\nSus medidas: el bebé mide alrededor de 20 cm de la coronilla a las nalgas y pesa casi medio kilo.'),
(24, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/24SM.jpg', 'Con 24 semanas de embarazo, los primeros estímulos del exterior van llegando al feto. Casi todos sus pequeños órganos sensoriales (oído, olfato, papilas gustativas y nervios del tacto) están ya maduros. Esta etapa es clave porque tu bebé empieza a interpretar el mundo, a interactuar, explorar, aprender. Para empezar, se va familiarizando con olores y sabores del exterior y de la propia madre (como los de la leche) a través del líquido amniótico. Si le gustan, esto lo animará a comer cuando nazca.\r\nEl único sentido que tu pequeño todavía no experimenta durante estas semanas es el de la vista. Puede percibir algún brillo  de una luz fuerte como la del sol, pero el útero tiene las paredes muy gruesas y es muy oscuro. Aun así, los bebés abren y cierran los ojos en esta etapa. Este movimiento es el precursor del reflejo del parpadeo.\r\nSus medidas: supera el medio kilo y va a crecer hasta los 22 centímetros.'),
(25, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/25SM.jpg', 'En la semana 25 de embarazo el bebé ya tiene pestañas, aunque el color de los ojos todavía no se ha desarrollado por completo, ya que algunos pigmentos necesitan luz para acabar de formarse. Por eso, los ojos de tu bebé pueden cambiar en las primeras semanas de vida.\r\nLos asiáticos y los africanos nacen con ojitos marrones o grises y acaban teniéndolos oscuros o negros. Los caucásicos suelen nacer con ojos claros, pero cuando maduran no tienen porqué conservarlos, posiblemente acaben con ojos verdes o marrones. Antes creíamos que el color de los ojos estaba determinado por un solo gen, nuevas investigaciones nos han demostrado que realmente son varios genes, así que es imposible saber qué ojitos tendrá tu bebé sólo mirando el de los progenitores.\r\nEl sentido que más se desarrolla a partir de esta semana es el oído. El niño está muy aislado, pero las ondas sonoras viajan más rápido por el líquido amniótico que por el aire. Como consecuencia, tu pequeño empieza a escuchar los primeros sonidos, principalmente tus gorgoteos y los murmullos de tu cuerpo. También percibirá ruidos que él mismo produce, como sus chapoteos en el líquido amniótico o el movimiento de los líquidos producido por las ondas ultrasónicas de las ecografías. Aunque los ultrasonidos no son percibidos por el oído humano, las ondas agitan el fluido de la bolsa amniótica y el pequeño puede percibir su sonido. Del mundo exterior, podrán llegarle conversaciones, ruidos estridentes y música.  Con todo, los soniditos de la mamá siempre serán diferentes del resto, porque viajan a través de los fluidos de los dos cuerpos. Ésta es una de las muchas explicaciones que se da a la especial relación que establecen madres e hijos desde el nacimiento.\r\nSus medidas: tu bebé mide entorno a los 22 centímetros y empieza a acercarse al kilo.'),
(26, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/26SM.jpg', 'Tu pequeñín empezará a hacer uso de más reflejos importantes como el reflejo labial de succión, ya que empieza a chuparse el dedo con virulencia. La mayor parte del día, tu bebé estará durmiendo. Así que el tiempo que pase despierto tendrá más energía y, por tanto, estará muy alerta a todo lo que sucede en su entorno.\r\nCuando oiga ruidos repentinos, sacará sus reflejos de protección extendiendo sus brazos y piernas como mecanismo de autodefensa. Hoy en día, con las comodidades de la vida humana, este instinto no se hace tan necesario como lo era para nuestros antepasados. Sin embargo, otros reflejos sí lo son. Es el caso del reflejo que consiste en ingerir medio litro de líquido amniótico al día. Esta tendencia natural ayuda a que el sistema digestivo se desarrolle bien.\r\nSus medidas: 23 centímetros y un kilo de peso aproximadamente.'),
(27, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/27SM.jpg', 'Ya has llegado a la semana 27 de embarazo. En este momento,  ya se podrá escuchar los latiditos del bebé poniendo los oídos sobre tu abdomen. Debes saber que tu ritmo cardíaco está muy vinculado al de tu hijo y tu estrés y tus hábitos le influyen directamente. Si no llevas unas costumbres y un ritmo de vida sanos, puede desarrollar diversos problemas.\r\nEstas semanas notarás movimientos a diario. Existe incluso la posibilidad de que percibas momentos en los que tu bebé tiene hipo. El hipo del feto es muy curioso, muy diferente al nuestro: tiene espasmos, pero no produce ruido porque no hay aire en sus pulmones. Entrañable.\r\nPrecisamente los pulmones son el último órgano vital que se forma en el bebé. En tu interior no los usa, ya que obtiene el oxígeno de tu placenta a través del cordón umbilical y también de lo que traga de la bolsa amniótica. Sin embargo, los pequeños músculos de su pecho empiezan a practicar un movimiento como el de la respiración empleando los pulmones y el diafragma. ¡Se va preparando para la vida aquí afuera!\r\nSus medidas: 24 centímetros y 1 kilo de peso aproximadamente.'),
(28, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/28SM.jpg', 'Tu bebé sigue aumentando de peso. Ya han pasado 28 semanas de embarazo y es hora de que empiece a producir una capa grasa bajo la piel.\r\nAdemás, sus sentidos están cada vez más despiertos y su corteza cerebral se ha desarrollado hasta tal punto que puede empezar a albergar conocimiento!. Es más, en las próximas semanas, su sistema nervioso estará tan avanzado como el del recién nacido. Investigaciones recientes explican que, con la semana 28, el pequeño es más consciente de lo que le rodea. Una semana de éstas, creará su primer recuerdo.\r\nTe encantará saber que, en este momento del embarazo, tu pequeño empieza a familiarizarse con tu voz. La reconoce, incluso responde a ella, como lo puede hacer con la música. Hay estudios que confirman que si escucha una misma melodía una y otra vez también puede reconocerla, incluso seguir su ritmo.\r\nEsto ha demostrado que la memoria a largo plazo funciona antes del nacimiento. Impresionante ¿verdad? Debes tener en cuenta que las canciones pueden influir en su ánimo. Las de ritmos acelerados lo sobrestimularán, en cambio, la música suave le relajará.\r\nSus medidas: tu bebé ya habrá crecido hasta aproximadamente los 37 centímetros y empezará a superar el kilo.'),
(29, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/29SM.jpg', 'El feto ya está muy desarrollado: los sentidos están muy activos, el niño ya sabe ponerse cómodo y ha aprendido a moverse. A partir de las 29 semanas de embarazo, empezará a darse la vuelta para reposar hacia abajo, apoyado en el cuello de tu útero. Sus idas y venidas harán que te dé más de una patadita en las costillas. Lo notarás.\r\nA partir de la semana 29 entramos en una fase en la que su cerebro ha madurado tanto que puede regular su temperatura corporal. Por supuesto, el bebé todavía necesita el calor del cuerpo de su madre para mantenerse caliente hasta el que nazca.\r\nSe sigue desarrollando células nerviosas del cerebro. Al nacer, tendrá cientos de miles de millones de ellas, que aunque parezcan muchas, se debe recordar que no se crearán más después de nacer.\r\nSus medidas: el bebé pesará alrededor de 1kilo con 400 gramos y ya alcanza los 42 centímetros.'),
(30, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/30SM.jpg', 'El pequeño sigue incrementando su peso, en parte porque agrega nuevas capas de vérnix, la grasita que le ayudará a regular su temperatura corporal y le mantendrá abrigado incluso después de nacer. En la semana 30 de embarazo, la piel del bebé no será tan arrugadita, estará más tersa.\r\nAdemás ya comienza a buscar la posición definitiva que tendrá al nacer. La mayoría de los bebés se ubican con la cabeza hacia abajo, pero algunos deciden no hacerlo, ya sea porque se acuerdan tarde de girar y el útero ya no lo permite, o bien porque se hallan enredados con el cordón umbilical o el mismo resulta corto, limitando su movilidad.\r\nSus medida: en la semana 30 pesará aproximadamente un kilo y medio y medirá unos 43 centímetros .'),
(31, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/31SM.jpg', 'En la semana 31 de embarazo, al bebé le cuesta moverse y estirarse dentro del útero cada vez más porque su crecimiento se acelera. Si no ha acomodado todavía su cabeza hacia abajo, lo hará cerca de esta semana.\r\nLos pulmones están desarrollados casi al completo. Su actividad y su rápido crecimiento continúan en todos los aspectos. Sin ir más lejos, cada día tu bebé elimina aproximadamente medio litro de orina al líquido amniótico.\r\nSus medidas: el pequeño se acerca a la longitud que tendrá en el nacimiento. Ya son 31 semanas de gestación y pesa más de kilo y medio.'),
(32, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/32SM.jpg', 'Hoy en día los científicos ven pocas diferencias entre el feto de 32 semanas de embarazo y el cerebro del recién nacido. Antes se pensaba que el desarrollo mental empezaba con el nacimiento. Pero ahora se cree que el bebé dentro del útero puede pensar, incluso hacer memoria. A partir de esta semana 32, puede incluso crear su primer recuerdo. Las uñas llegan a las puntas de los dedos de la mano, así que, aunque te sorprenda, deberás cortárselas al poco tiempo de nacer.\r\nSus medidas: en la semana 32 de embarazo tu bebe medirá unos 47 centímetros y  pesará casi 2 kilos.'),
(33, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/33SM.jpg', 'Se ha descubierto que a partir de la semana 33 de embarazo, el feto realiza unos Movimientos Oculares Rápidos (MOR) que pueden ser señal de que el niño sueña.\nCon las pocas experiencias vitales que tiene un bebé dentro de su madre, cuesta imaginar con qué pueden soñar estos renacuajos. ¿Que juegan con sus piecitos? ¿Con tus ruiditos? Para un feto, soñar, a pesar de la simplicidad de estas ensoñaciones, puede jugar papel crucial en cuanto a la estimulación y el crecimiento del cerebro.\n\nEn la semana 33, el cerebro del bebé prosigue su rápido desarrollo, teniendo los cinco sentidos ya en funcionamiento. Puede ver el mundo líquido que le rodea, saborear el líquido amniótico que traga, sentir el tacto del dedo que chupa, y oír el corazón y la voz de su madre. En el saco amniótico no hay aire que le lleve olores, pero si lo hubiera, también podría olerlos. Debido al gran desarrollo cerebral experimentado en esta semana 33, la circunferencia craneal del bebé ha aumentado en los últimos días casi 1,25 centímetros. Sus medidas: en esta semana 33, el bebé ya está hecho un grandullón o una grandullona. Mide unos 47 centímetros y sobrepasa los 2 kilos de peso.'),
(34, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/34SM.jpg', 'En la semana 34 del embarazo, las conexiones cerebrales van a un ritmo trepidante y la cabecita de tu bebé va creciendo con ellas.\r\nEl bebé percibe estímulos, a los que reacciona más vivo que nunca. Aunque el bebé está en una etapa en la que duerme mucho, está muy atento a todo lo que le rodea y cualquier cosa que le quite el sueño puede afectarle. Si esto ocurre de forma continuada, puede ser perjudicial. Así que recuerda que tu embarazo es muy importante y adapta tu ritmo de vida a tu estado.\r\nYa estás en la semana 34, sé paciente. ¡Ya queda poco para dar luz! \r\nSus medidas: el bebé está cerca de los 50 centímetros y su cuerpecito ya pesa 2 kilos.'),
(35, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/35SM.jpg', 'Ya son 8 meses. Tu pequeño ya es un bebé apretadito y pesado. Y es que su cerebro y su cabeza ya han alcanzado su tamaño máximo. Tu hijo ha producido 100 billones neuronas con 100 trillones de conexiones, que le durarán toda la vida. ¡Y empezó hace 28 semanas! (en la semana 7).\r\nSi tuvieras un parto prematuro a estas alturas, podría sobrevivir sin problemas. Aunque, claro está, cuanto más tiempo esté dentro de tu vientre, más desarrollado y sano estará en el momento de nacer.\r\nEn esta semana 35 y en las últimas semanas, tu cuerpo le transferirá a tu hijo inmunidad temporal contra enfermedades infantiles (como las paperas y el sarampión). El bebé estará protegido hasta que le pongas las primeras vacunas. Su piel se alisa y el lanugo comienza a caer.\r\nSus medidas: a partir de esta semana 35 comienza el período de aumento de peso más rápido del bebé. Ganará entre 250 y 350 gramos por semana. Ya está cerca de los 3 kilos.'),
(36, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/36SM.jpg', 'Es posible que notes menos movimiento sobre la semana 36 del embarazo. Esto ocurre porque el bebé ha crecido tanto que tiene menos espacio para moverse. Asimismo, se le forman cúmulos de grasa que le redondean el cuerpo.\r\nPor otro lado, la piel empieza a hacerse más rosada.\r\nEn cualquier momento se encaja del todo y acaba de poner su cabecita en el cuello del útero. Puede superar los 50 centímetros, ¡qué grandote!'),
(37, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/37SM.jpg', 'Estás en la semana 37 y a partir de ahora, el bebé está en condiciones óptimas para nacer. En esta semana 37 ya habrá adoptado la posición definitiva. Lo más probable es que esté cabeza abajo y con la cara mirando hacia atrás, listo para nacer.\r\nEl médico te podrá decir a través de una ecografía si se encuentra en otra posición que haga necesaria una cesáreapara evitar riesgos (cabeza girada para hacia delante o de nalgas).\r\nEn la semana 37 el cerebro y el cráneo del bebé también continúan creciendo. Desde la semana 37 no va a aumentar mucho más de peso, a pesar de ello estas semanas siguen siendo importantes, ya que todavía acumula 15 gramos de grasa al día. Con esta grasa su cuerpo regulará mejor la temperatura, manteniendo un nivel adecuado de azúcar en sangre.\r\nSus medidas: en la semana 37 la mayoría de bebés suelen medir unos 50 centímetros de largo, con un peso de entre 2,7 y 3 kg.'),
(38, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/38SM.jpg', 'Si tu parto no se adelanta, con la semana 38 de embarazo empieza la cuenta atrás. Es imposible predecir cuándo nacerá exactamente el bebé. De hecho, sólo el 5% nace en la fecha prevista, el resto lo hace antes o después.\r\nEn la semana 38 debes estar alerta a las señales que te indiquen que has roto aguas y controlar las contracciones, que en las últimas semanas habrán sido numerosas, aunque separadas. Atenta: si son muy intensas y continuadas (más de 5 por hora) es posible que estés de parto.\r\nNadie sabe qué desencadena el parto exactamente. Lo que sí sabemos es que los pulmones del niño y tu placenta son la clave de la sincronización. Cuando los pulmones están maduros, segregan una proteína al líquido amniótico que altera la producción hormonas. Este cambio hace que la placenta reduzca la emisión de la hormona progesterona y fomente la producción de una nueva hormona: la oxitocina.\r\nLa oxitocina regula las contracciones del útero e indica si hay parto. También bloquea tus recuerdos y te ayuda a olvidar el dolor y unirte al bebé.'),
(39, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/39SM.jpg', 'Estás en la semana 39, el bebé traga líquido amniótico y empieza a acumularlo como material de desecho, que se denomina meconio. El meconio es una sustancia negra pegajosa que será su primer movimiento de intestinos después del nacimiento. En definitiva, su primera caca.\r\nEl cordón umbilical, que hasta el momento ha trasportado los nutrientes desde la placenta al bebé, mide, en esta semana 39  unos 50 centímetros de largo y 1,3 centímetros de ancho. Como el bebé ocupa todo el espacio en el útero, es común que el cordón umbilical se enrolle a su alrededor. Éste es muy elástico y no suele dar problemas. Aun así, hay casos en los que el bebé pueda nacer con el cordón rodeando su cuello. Tranquila, los partos con más de 35 semanas hoy en día son fáciles, y cortar el cordón umbilical también.\r\nEstás en la semana 39, tu parto se puede iniciar en cualquier momento.'),
(40, 'https://s3.eu-central-1.wasabisys.com/stela/weeks/40SM.jpg', 'En la semana 40 el feto tiene el tamaño completo y está listo para nacer. La mayor parte de vérmix (grasa que lo cubre) ha desaparecido, aunque pueden quedar algunos restos en sus pliegues. Ya tendrá pelito y uñas largas.\r\nNo te preocupes si aún no ha nacido, el bebé está preparado y es cuestión de tiempo. Los controles médicos serán continuos para asegurar que no haya ningún problema. Si tienes dudas consulta con tu médico. Es la semana 40, descansa y revisa que tienes preparado todo lo que necesitas.\r\nNadie sabe qué desencadena el parto exactamente. Lo que sí sabemos es que los pulmones del niño y tu placenta son la clave de la sincronización. Cuando los pulmones están maduros, segregan una proteína al líquido amniótico que altera la producción hormonas. Este cambio hace que la placenta reduzca la emisión de progesterona y fomente la producción de una nueva hormona: la oxitocina. La oxitocina regula las contracciones del útero e indica si hay parto. También bloquea tus recuerdos y te ayuda a olvidar el dolor y unirte al bebé. Es la semana 40, puedes dar a luz en cualquier momento. Ve al hospital, llama a tu médico o prepara todo para tener el parto que hayas planeado (convencional, natural, en el agua…).\r\nSus medidas: la longitud de tu bebé varía entre los 48 y los 53 cm y su peso puede estar entre los 3 y 4,5 kg.');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `clinics`
--

CREATE TABLE `clinics` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL,
  `address` varchar(64) DEFAULT NULL,
  `available` int(1) NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `clinics`
--

INSERT INTO `clinics` (`id`, `user_id`, `name`, `address`, `available`, `updated_at`, `created_at`) VALUES
(1, 1, 'Master', ' ', 1, '2020-11-25 13:25:26', '2020-11-23 00:00:00'),
(10, 2, 'Baby&Me', ' ', 1, '2020-11-25 13:25:26', '2020-11-23 00:00:00');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `heartbeats`
--

CREATE TABLE `heartbeats` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `url` varchar(256) NOT NULL,
  `size` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `multimedia_contents`
--

CREATE TABLE `multimedia_contents` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `type` varchar(32) NOT NULL,
  `url` varchar(255) NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime NOT NULL DEFAULT '2000-01-01 01:01:01',
  `size` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `multimedia_contents`
--

INSERT INTO `multimedia_contents` (`id`, `user_id`, `name`, `type`, `url`, `updated_at`, `created_at`, `deleted_at`, `size`) VALUES
(539, 4, 'MARIANA VICENTE (17).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2817%29.png-compress.jpg', '2021-02-23 11:01:44', '2021-02-23 11:01:44', '2000-01-01 01:01:01', 452684),
(540, 4, 'MARIANA VICENTE (18).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2818%29.png-compress.jpg', '2021-02-23 11:01:44', '2021-02-23 11:01:44', '2000-01-01 01:01:01', 452452),
(549, 4, 'MARIANA VICENTE (2).mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29.mp4-audio.mp4', '2021-02-23 11:02:09', '2021-02-23 11:02:09', '2000-01-01 01:01:01', 10940266),
(550, 4, 'MARIANA VICENTE (13).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2813%29.png-compress.jpg', '2021-02-23 11:27:00', '2021-02-23 11:27:00', '2000-01-01 01:01:01', 502598),
(551, 4, 'MARIANA VICENTE (8).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%288%29.png-compress.jpg', '2021-02-23 11:28:16', '2021-02-23 11:28:16', '2000-01-01 01:01:01', 494218),
(552, 4, 'MARIANA VICENTE (19).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2819%29.png-compress.jpg', '2021-02-23 11:28:23', '2021-02-23 11:28:23', '2000-01-01 01:01:01', 442226),
(553, 4, 'MARIANA VICENTE (17).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2817%29.png-compress.jpg', '2021-02-23 11:28:24', '2021-02-23 11:28:24', '2000-01-01 01:01:01', 452684),
(554, 4, 'MARIANA VICENTE (22).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2822%29.png-compress.jpg', '2021-02-23 11:28:24', '2021-02-23 11:28:24', '2000-01-01 01:01:01', 422193),
(555, 4, 'MARIANA VICENTE (18).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2818%29.png-compress.jpg', '2021-02-23 11:28:24', '2021-02-23 11:28:24', '2000-01-01 01:01:01', 452452),
(556, 4, 'MARIANA VICENTE (16).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2816%29.png-compress.jpg', '2021-02-23 11:28:24', '2021-02-23 11:28:24', '2000-01-01 01:01:01', 491212),
(557, 4, 'MARIANA VICENTE (20).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2820%29.png-compress.jpg', '2021-02-23 11:28:24', '2021-02-23 11:28:24', '2000-01-01 01:01:01', 434796),
(558, 4, 'MARIANA VICENTE (15).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2815%29.png-compress.jpg', '2021-02-23 11:28:24', '2021-02-23 11:28:24', '2000-01-01 01:01:01', 495652),
(559, 4, 'MARIANA VICENTE (21).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2821%29.png-compress.jpg', '2021-02-23 11:28:25', '2021-02-23 11:28:25', '2000-01-01 01:01:01', 450378),
(560, 4, 'MARIANA VICENTE (6).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%286%29.png-compress.jpg', '2021-02-23 11:30:49', '2021-02-23 11:30:49', '2000-01-01 01:01:01', 470956),
(561, 4, 'MARIANA VICENTE (11).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2811%29.png-compress.jpg', '2021-02-23 11:52:16', '2021-02-23 11:52:16', '2000-01-01 01:01:01', 495473),
(562, 4, 'MARIANA VICENTE (11).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2811%29.png-compress.jpg', '2021-02-23 11:53:05', '2021-02-23 11:53:05', '2000-01-01 01:01:01', 495473),
(563, 4, 'MARIANA VICENTE (11).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2811%29.png-compress.jpg', '2021-02-23 11:53:49', '2021-02-23 11:53:49', '2000-01-01 01:01:01', 495473),
(564, 4, 'MARIANA VICENTE (4).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%284%29.png-compress.jpg', '2021-02-23 11:56:08', '2021-02-23 11:56:08', '2000-01-01 01:01:01', 447875),
(565, 4, 'MARIANA VICENTE (11).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2811%29.png-compress.jpg', '2021-02-23 12:01:17', '2021-02-23 12:01:17', '2000-01-01 01:01:01', 495473),
(566, 4, 'MARIANA VICENTE (13).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2813%29.png-compress.jpg', '2021-02-23 12:05:40', '2021-02-23 12:05:40', '2000-01-01 01:01:01', 502598),
(567, 4, 'MARIANA VICENTE (14).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2814%29.png-compress.jpg', '2021-02-23 12:07:08', '2021-02-23 12:07:08', '2000-01-01 01:01:01', 478726),
(568, 4, 'MARIANA VICENTE (8).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%288%29.png-compress.jpg', '2021-02-23 12:11:00', '2021-02-23 12:11:00', '2000-01-01 01:01:01', 494218),
(569, 4, 'MARIANA VICENTE (13).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2813%29.png-compress.jpg', '2021-02-23 12:13:48', '2021-02-23 12:13:48', '2000-01-01 01:01:01', 502598),
(570, 4, 'MARIANA VICENTE (11).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2811%29.png-compress.jpg', '2021-02-23 12:15:53', '2021-02-23 12:15:53', '2000-01-01 01:01:01', 495473),
(571, 4, 'MARIANA VICENTE (11).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2811%29.png-compress.jpg', '2021-02-23 12:24:39', '2021-02-23 12:24:39', '2000-01-01 01:01:01', 495473),
(572, 4, 'MARIANA VICENTE (2).mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29.mp4-audio.mp4', '2021-02-23 12:31:12', '2021-02-23 12:31:12', '2000-01-01 01:01:01', 10940266),
(573, 4, 'MARIANA VICENTE (1).mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%281%29.mp4-audio.mp4', '2021-02-23 12:38:14', '2021-02-23 12:38:14', '2000-01-01 01:01:01', 17452045),
(574, 4, 'MARIANA VICENTE (11).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2811%29.png-compress.jpg', '2021-02-23 12:39:30', '2021-02-23 12:39:30', '2000-01-01 01:01:01', 495473),
(575, 4, 'MARIANA VICENTE (7).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%287%29.png-compress.jpg', '2021-02-23 12:39:37', '2021-02-23 12:39:37', '2000-01-01 01:01:01', 505559),
(576, 4, 'MARIANA VICENTE (2) (online-video-cutter.com).mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4', '2021-02-23 12:43:12', '2021-02-23 12:43:12', '2000-01-01 01:01:01', 821529),
(577, 4, 'MARIANA VICENTE (14).png-compress.jpg-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2814%29.png-compress.jpg-compress.jpg', '2021-02-23 12:47:14', '2021-02-23 12:47:14', '2000-01-01 01:01:01', 478726),
(578, 4, 'MARIANA VICENTE (100).png-compress.jpg-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%28100%29.png-compress.jpg-compress.jpg', '2021-02-23 12:47:20', '2021-02-23 12:47:20', '2000-01-01 01:01:01', 347459),
(579, 4, 'MARIANA VICENTE (2).mp4-audio.mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29.mp4-audio.mp4-audio.mp4', '2021-02-23 12:48:27', '2021-02-23 12:48:27', '2000-01-01 01:01:01', 9800998),
(580, 4, 'MARIANA VICENTE (2) (online-video-cutter.com).mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4', '2021-02-23 12:58:45', '2021-02-23 12:58:45', '2000-01-01 01:01:01', 821529),
(581, 4, 'MARIANA VICENTE (4).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%284%29.png-compress.jpg', '2021-02-23 14:30:43', '2021-02-23 14:30:43', '2000-01-01 01:01:01', 447875),
(582, 4, 'MARIANA VICENTE (4).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%284%29.png-compress.jpg', '2021-02-23 14:32:55', '2021-02-23 14:32:55', '2000-01-01 01:01:01', 447875),
(583, 4, 'MARIANA VICENTE (2) (online-video-cutter.com).mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4', '2021-02-23 14:57:21', '2021-02-23 14:57:21', '2000-01-01 01:01:01', 821529),
(584, 4, 'MARIANA VICENTE (2).mp4-audio.mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29.mp4-audio.mp4-audio.mp4', '2021-02-23 14:58:19', '2021-02-23 14:58:19', '2000-01-01 01:01:01', 9800998),
(585, 4, 'MARIANA VICENTE (5).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%285%29.png-compress.jpg', '2021-02-23 15:14:23', '2021-02-23 15:14:23', '2000-01-01 01:01:01', 498982),
(586, 4, 'MARIANA VICENTE (100).png-compress.jpg-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%28100%29.png-compress.jpg-compress.jpg', '2021-02-23 15:15:58', '2021-02-23 15:15:58', '2000-01-01 01:01:01', 347459),
(587, 4, 'MARIANA VICENTE (100).png-compress.jpg-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%28100%29.png-compress.jpg-compress.jpg', '2021-02-23 15:18:32', '2021-02-23 15:18:32', '2000-01-01 01:01:01', 347459),
(588, 4, 'MARIANA VICENTE (100).png-compress.jpg-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%28100%29.png-compress.jpg-compress.jpg', '2021-02-23 15:19:32', '2021-02-23 15:19:32', '2000-01-01 01:01:01', 347459),
(589, 4, 'MARIANA VICENTE (100).png-compress.jpg-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%28100%29.png-compress.jpg-compress.jpg', '2021-02-23 15:20:19', '2021-02-23 15:20:19', '2000-01-01 01:01:01', 347459),
(590, 4, 'MARIANA VICENTE.png-compress.jpg-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE.png-compress.jpg-compress.jpg', '2021-02-23 15:24:17', '2021-02-23 15:24:17', '2000-01-01 01:01:01', 505960),
(591, 4, 'MARIANA VICENTE (2) (online-video-cutter.com).mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4', '2021-02-23 15:27:45', '2021-02-23 15:27:45', '2000-01-01 01:01:01', 821529),
(592, 4, 'MARIANA VICENTE (2) (online-video-cutter.com).mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4', '2021-02-23 15:29:29', '2021-02-23 15:29:29', '2000-01-01 01:01:01', 821529),
(593, 4, 'MARIANA VICENTE (2) (online-video-cutter.com).mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4', '2021-02-23 15:31:34', '2021-02-23 15:31:34', '2000-01-01 01:01:01', 821529),
(594, 4, 'MARIANA VICENTE (2).mp4-audio.mp4-audio.mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29.mp4-audio.mp4-audio.mp4-audio.mp4', '2021-02-23 15:43:16', '2021-02-23 15:43:16', '2000-01-01 01:01:01', 8957933),
(595, 4, 'MARIANA VICENTE (2) (online-video-cutter.com).mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4', '2021-02-23 15:45:10', '2021-02-23 15:45:10', '2000-01-01 01:01:01', 821529),
(596, 4, 'MARIANA VICENTE (14).png-compress.jpg-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2814%29.png-compress.jpg-compress.jpg', '2021-02-23 15:48:28', '2021-02-23 15:48:28', '2000-01-01 01:01:01', 478726),
(597, 4, 'MARIANA VICENTE (2) (online-video-cutter.com).mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4', '2021-02-23 15:50:39', '2021-02-23 15:50:39', '2000-01-01 01:01:01', 821529),
(598, 4, 'MARIANA VICENTE (2).mp4-audio.mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29.mp4-audio.mp4-audio.mp4', '2021-02-23 15:50:46', '2021-02-23 15:50:46', '2000-01-01 01:01:01', 9800998),
(599, 4, 'MARIANA VICENTE.png-compress.jpg-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE.png-compress.jpg-compress.jpg', '2021-02-23 15:53:28', '2021-02-23 15:53:28', '2000-01-01 01:01:01', 505960),
(600, 4, 'MARIANA VICENTE (100).png-compress.jpg-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%28100%29.png-compress.jpg-compress.jpg', '2021-02-23 15:53:38', '2021-02-23 15:53:38', '2000-01-01 01:01:01', 347459),
(601, 4, 'MARIANA VICENTE.png-compress.jpg-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE.png-compress.jpg-compress.jpg', '2021-02-23 15:53:49', '2021-02-23 15:53:49', '2000-01-01 01:01:01', 505960),
(602, 4, 'MARIANA VICENTE (2) (online-video-cutter.com).mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4', '2021-02-23 15:55:00', '2021-02-23 15:55:00', '2000-01-01 01:01:01', 821529),
(603, 4, 'MARIANA VICENTE (2) (online-video-cutter.com).mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4', '2021-02-23 15:59:39', '2021-02-23 15:59:39', '2000-01-01 01:01:01', 821529),
(604, 4, 'MARIANA VICENTE (2).mp4-audio.mp4-audio.mp4', 'video', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29.mp4-audio.mp4-audio.mp4', '2021-02-23 16:03:35', '2021-02-23 16:03:35', '2000-01-01 01:01:01', 9800998),
(605, 4, 'MARIANA VICENTE (3).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%283%29.png-compress.jpg', '2021-02-24 16:56:13', '2021-02-24 16:56:13', '2000-01-01 01:01:01', 441803),
(606, 4, 'MARIANA VICENTE (4).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%284%29.png-compress.jpg', '2021-02-24 16:56:14', '2021-02-24 16:56:14', '2000-01-01 01:01:01', 447875),
(607, 4, 'MARIANA VICENTE (6).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%286%29.png-compress.jpg', '2021-02-24 16:56:14', '2021-02-24 16:56:14', '2000-01-01 01:01:01', 470956),
(608, 4, 'MARIANA VICENTE (9).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%289%29.png-compress.jpg', '2021-02-24 16:56:14', '2021-02-24 16:56:14', '2000-01-01 01:01:01', 502197),
(609, 4, 'MARIANA VICENTE (5).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%285%29.png-compress.jpg', '2021-02-24 16:56:14', '2021-02-24 16:56:14', '2000-01-01 01:01:01', 498982),
(610, 4, 'MARIANA VICENTE (7).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%287%29.png-compress.jpg', '2021-02-24 16:56:14', '2021-02-24 16:56:14', '2000-01-01 01:01:01', 505559),
(611, 4, 'MARIANA VICENTE (8).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%288%29.png-compress.jpg', '2021-02-24 16:56:14', '2021-02-24 16:56:14', '2000-01-01 01:01:01', 494218),
(612, 4, 'MARIANA VICENTE (11).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2811%29.png-compress.jpg', '2021-02-24 16:56:15', '2021-02-24 16:56:15', '2000-01-01 01:01:01', 495473),
(613, 4, 'MARIANA VICENTE (13).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2813%29.png-compress.jpg', '2021-02-24 16:56:15', '2021-02-24 16:56:15', '2000-01-01 01:01:01', 502598),
(614, 4, 'MARIANA VICENTE (10).png-compress.jpg', 'image', 'https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA%20VICENTE%20%2810%29.png-compress.jpg', '2021-02-24 16:56:15', '2021-02-24 16:56:15', '2000-01-01 01:01:01', 502909);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `streamings`
--

CREATE TABLE `streamings` (
  `id` int(11) NOT NULL,
  `clinic_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `url` varchar(64) NOT NULL,
  `code` varchar(4) NOT NULL,
  `available` int(1) NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `streamings`
--

INSERT INTO `streamings` (`id`, `clinic_id`, `user_id`, `url`, `code`, `available`, `updated_at`, `created_at`) VALUES
(10, 10, 4, 'https://www.youtube.com/watch?v=TC3BhhraHgc', 'YEHG', 1, '2021-02-24 16:57:46', '2021-02-24 16:57:46');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `clinic_id` int(11) DEFAULT NULL,
  `username` varchar(64) NOT NULL,
  `password` varchar(256) NOT NULL,
  `name` varchar(64) DEFAULT '',
  `lastname` varchar(64) DEFAULT '',
  `phone` varchar(15) DEFAULT '',
  `rol` varchar(20) NOT NULL,
  `available` int(1) NOT NULL,
  `token` varchar(1024) DEFAULT '',
  `firebase_token` varchar(1024) DEFAULT '',
  `push_token` varchar(255) DEFAULT '',
  `device_type` varchar(60) DEFAULT '',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `users`
--

INSERT INTO `users` (`id`, `clinic_id`, `username`, `password`, `name`, `lastname`, `phone`, `rol`, `available`, `token`, `firebase_token`, `push_token`, `device_type`, `updated_at`, `created_at`) VALUES
(1, 1, 'pauloxti@gmail.com', '$2y$04$EYx4ddk6C5W/Qu2x7QozP.03/tdxCzP/PvACjNMPG0RFKve47wv9y', 'Master', 'Master', '666666666', 'Master', 1, '', '', '', 'web', '2021-01-18 12:02:05', '2020-01-14 14:54:37'),
(2, 10, 'info@babyandme.pt', '$2y$04$EYx4ddk6C5W/Qu2x7QozP.03/tdxCzP/PvACjNMPG0RFKve47wv9y', 'BabyAndMe', 'BabyAndMe', '', 'Propietario', 1, '', '', '', 'web', '2021-02-24 18:11:37', '2020-01-14 14:54:37'),
(4, 10, 'test@gmail.com', '$2a$04$VCKfcG7kcQbhgJKPN0C66e.xteAY7/5DeiPTGzv1RAuv.TYtYuxQq', 'Cliente', 'Uno', '666666666', 'Cliente', 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InRlc3RAZ21haWwuY29tIiwiZXhwIjoxNjIwMDIxOTg1fQ.VAKFVbYyXwprAbdq1aaAmwEcVNua-fFDXD6RNBYEA90', '', '', 'ios', '2021-02-22 16:46:25', '2020-01-14 14:54:37');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `users_promos`
--

CREATE TABLE `users_promos` (
  `emitter_user_id` int(11) NOT NULL,
  `title` varchar(64) NOT NULL,
  `text` varchar(255) NOT NULL,
  `available` int(1) NOT NULL DEFAULT '1',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `start_at` datetime NOT NULL,
  `end_at` datetime DEFAULT NULL,
  `week` int(11) NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `users_promos`
--

INSERT INTO `users_promos` (`emitter_user_id`, `title`, `text`, `available`, `updated_at`, `created_at`, `start_at`, `end_at`, `week`) VALUES
(2, 'Promo de prueba numero 1', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore', 1, '2020-08-05 18:13:01', '2012-11-15 00:00:00', '2020-04-10 00:00:00', '2021-02-26 00:00:00', 12),
(2, 'Promo de prueba numero 2', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore', 1, '2020-08-05 18:13:01', '2019-11-15 00:00:00', '2020-04-10 00:00:00', '2020-12-17 00:00:00', 12),
(2, 'Promo de prueba numero 3', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore', 1, '2020-08-05 18:13:01', '2019-11-15 00:00:01', '2020-04-10 00:00:00', '2020-12-25 00:00:00', 12);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `users_recovery`
--

CREATE TABLE `users_recovery` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `token` varchar(255) NOT NULL,
  `available` int(1) NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `video_thumbnails`
--

CREATE TABLE `video_thumbnails` (
  `id` int(10) UNSIGNED NOT NULL,
  `video_id` int(10) UNSIGNED NOT NULL,
  `thumbnail` varchar(512) NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime NOT NULL DEFAULT '1000-01-01 00:00:00',
  `size` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `video_thumbnails`
--

INSERT INTO `video_thumbnails` (`id`, `video_id`, `thumbnail`, `updated_at`, `created_at`, `deleted_at`, `size`) VALUES
(79, 536, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 10:43:22', '2021-02-23 10:43:22', '1000-01-01 00:00:00', 56),
(80, 549, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 11:02:11', '2021-02-23 11:02:11', '1000-01-01 00:00:00', 62201),
(81, 572, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 12:31:15', '2021-02-23 12:31:15', '1000-01-01 00:00:00', 62201),
(82, 573, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%281%29.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 12:38:16', '2021-02-23 12:38:16', '1000-01-01 00:00:00', 58243),
(83, 576, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 12:43:15', '2021-02-23 12:43:15', '1000-01-01 00:00:00', 36278),
(84, 579, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29.mp4-audio.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 12:48:29', '2021-02-23 12:48:29', '1000-01-01 00:00:00', 58779),
(85, 580, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 12:58:47', '2021-02-23 12:58:47', '1000-01-01 00:00:00', 36278),
(86, 583, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 14:57:26', '2021-02-23 14:57:26', '1000-01-01 00:00:00', 36278),
(87, 584, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29.mp4-audio.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 14:58:21', '2021-02-23 14:58:21', '1000-01-01 00:00:00', 58779),
(88, 591, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 15:27:47', '2021-02-23 15:27:47', '1000-01-01 00:00:00', 36278),
(89, 592, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 15:29:32', '2021-02-23 15:29:32', '1000-01-01 00:00:00', 36278),
(90, 593, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 15:31:37', '2021-02-23 15:31:37', '1000-01-01 00:00:00', 36278),
(91, 594, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29.mp4-audio.mp4-audio.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 15:43:27', '2021-02-23 15:43:27', '1000-01-01 00:00:00', 57342),
(92, 595, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 15:45:22', '2021-02-23 15:45:22', '1000-01-01 00:00:00', 36278),
(93, 597, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 15:50:48', '2021-02-23 15:50:48', '1000-01-01 00:00:00', 36278),
(94, 598, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29.mp4-audio.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 15:50:53', '2021-02-23 15:50:53', '1000-01-01 00:00:00', 58779),
(95, 602, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 15:55:03', '2021-02-23 15:55:03', '1000-01-01 00:00:00', 36278),
(96, 603, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29%20%28online-video-cutter.com%29.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 15:59:43', '2021-02-23 15:59:43', '1000-01-01 00:00:00', 36278),
(97, 604, 'https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA%20VICENTE%20%282%29.mp4-audio.mp4-audio.mp4-thumbnail.jpg', '2021-02-23 16:03:40', '2021-02-23 16:03:40', '1000-01-01 00:00:00', 58779);

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `calculators`
--
ALTER TABLE `calculators`
  ADD PRIMARY KEY (`id`),
  ADD KEY `calculator_fk0` (`user_id`);

--
-- Indices de la tabla `calculator_details`
--
ALTER TABLE `calculator_details`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `clinics`
--
ALTER TABLE `clinics`
  ADD PRIMARY KEY (`id`),
  ADD KEY `clinics_fk0` (`user_id`);

--
-- Indices de la tabla `heartbeats`
--
ALTER TABLE `heartbeats`
  ADD PRIMARY KEY (`id`),
  ADD KEY `heartbeat_fk0` (`user_id`);

--
-- Indices de la tabla `multimedia_contents`
--
ALTER TABLE `multimedia_contents`
  ADD PRIMARY KEY (`id`),
  ADD KEY `multimedia_contents_fk0` (`user_id`);

--
-- Indices de la tabla `streamings`
--
ALTER TABLE `streamings`
  ADD PRIMARY KEY (`id`,`code`),
  ADD KEY `pk_clinics` (`clinic_id`),
  ADD KEY `pk_clinics_users` (`user_id`);

--
-- Indices de la tabla `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD KEY `users_fk0` (`clinic_id`);

--
-- Indices de la tabla `users_promos`
--
ALTER TABLE `users_promos`
  ADD PRIMARY KEY (`emitter_user_id`,`created_at`);

--
-- Indices de la tabla `users_recovery`
--
ALTER TABLE `users_recovery`
  ADD PRIMARY KEY (`id`),
  ADD KEY `users_recovery_fk0` (`user_id`);

--
-- Indices de la tabla `video_thumbnails`
--
ALTER TABLE `video_thumbnails`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `video_id` (`video_id`),
  ADD UNIQUE KEY `id` (`id`),
  ADD KEY `id_2` (`id`),
  ADD KEY `video_id_2` (`video_id`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `calculators`
--
ALTER TABLE `calculators`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=54;
--
-- AUTO_INCREMENT de la tabla `calculator_details`
--
ALTER TABLE `calculator_details`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=41;
--
-- AUTO_INCREMENT de la tabla `clinics`
--
ALTER TABLE `clinics`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
--
-- AUTO_INCREMENT de la tabla `heartbeats`
--
ALTER TABLE `heartbeats`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
--
-- AUTO_INCREMENT de la tabla `multimedia_contents`
--
ALTER TABLE `multimedia_contents`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=615;
--
-- AUTO_INCREMENT de la tabla `streamings`
--
ALTER TABLE `streamings`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
--
-- AUTO_INCREMENT de la tabla `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=70;
--
-- AUTO_INCREMENT de la tabla `users_recovery`
--
ALTER TABLE `users_recovery`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT de la tabla `video_thumbnails`
--
ALTER TABLE `video_thumbnails`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=98;
--
-- Restricciones para tablas volcadas
--

--
-- Filtros para la tabla `calculators`
--
ALTER TABLE `calculators`
  ADD CONSTRAINT `calculator_fk0` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Filtros para la tabla `clinics`
--
ALTER TABLE `clinics`
  ADD CONSTRAINT `clinics_fk0` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Filtros para la tabla `heartbeats`
--
ALTER TABLE `heartbeats`
  ADD CONSTRAINT `heartbeat_fk0` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Filtros para la tabla `multimedia_contents`
--
ALTER TABLE `multimedia_contents`
  ADD CONSTRAINT `multimedia_contents_fk0` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Filtros para la tabla `streamings`
--
ALTER TABLE `streamings`
  ADD CONSTRAINT `pk_clinics` FOREIGN KEY (`clinic_id`) REFERENCES `clinics` (`id`),
  ADD CONSTRAINT `pk_clinics_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Filtros para la tabla `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `users_fk0` FOREIGN KEY (`clinic_id`) REFERENCES `clinics` (`id`);

--
-- Filtros para la tabla `users_promos`
--
ALTER TABLE `users_promos`
  ADD CONSTRAINT `users_promos_fk0` FOREIGN KEY (`emitter_user_id`) REFERENCES `users` (`id`);

--
-- Filtros para la tabla `users_recovery`
--
ALTER TABLE `users_recovery`
  ADD CONSTRAINT `users_recovery_fk0` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
