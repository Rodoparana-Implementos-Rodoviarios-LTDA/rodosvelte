<?php
// carregar fontes...
$fnormal = '../font/arial.ttf';
$fnegrito = '../font/ARIBLK.ttf';
$fnegritoeitalico = '../font/Calibri.ttf';

$nome   = urldecode($_GET['nome']);
$cargo  = urldecode($_GET['cargo']);
$email  = urldecode($_GET['email']);
$skype  = urldecode($_GET['skype']);
$telefone    = urldecode($_GET['telefone']);
$celular    = urldecode($_GET['celular']);
$logo    = urldecode($_GET['logo']);
$filial = urldecode($_GET['filial']);
$ixa   = urldecode(isset($_GET['ixa']));
$ipo   = urldecode(isset($_GET['ipo']));
$ifu   = urldecode(isset($_GET['ifu']));
$ifl   = urldecode(isset($_GET['ifl']));
$isa   = urldecode(isset($_GET['isa']));
$eco   = urldecode(isset($_GET['eco']));
$pix   = urldecode(isset($_GET['pix']));


switch ($filial):
    case "curitiba":
        $endereco1 = "Av. Juscelino K. de Oliveira 3545";
        $endereco2 = "Cidade Industrial de Curitiba | Curitiba (PR)";
        break;
    case "lages":
        $endereco1 = "Av. Dr. JoÃ£o Pedro Arruda, 1437";
        $endereco2 = "Ãrea Industrial | Lages (SC)";
        break;
    case "guaiba":
        $endereco1 = "Rua da balanÃ§a, 96";
        $endereco2 = "Ramada | GuaÃ­ba (RS)";
        break;
    case "pelotas":
        $endereco1 = "BR 116, SN";
        $endereco2 = "Distrito RodoviÃ¡rio | CapÃ£o do LeÃ£o (RS)";
        break;
    case "treslagoas":
        $endereco1 = "BR 158, SN";
        $endereco2 = "Jardim Paranapunga | TrÃªs Lagoas (MS)";
        break;
    case "curvelo":
        $endereco1 = "Av. Bias Fortes, 2015";
        $endereco2 = "Tibira | Curvelo (MG)";
        break;
    case "palhoca":
        $endereco1 = "Rua JosÃ© Cosme Pamplona";
        $endereco2 = "Bela Vista | PalhoÃ§a (SC)";
        break;
endswitch;
1
?>
<style type="text/css">
    a:link {
        color: #000000;
    }

    p {
        margin: 2px
    }
</style>
<br>
<table border="0px">
    <tr rowspan="2">
        <td valign="top" class="p-2">
            <?php if ($logo == "ti") : ?>
                <img src="http://140.238.186.242/signaturegen/imagens/timber.png" width="100"></br></br>
            <?php elseif ($logo == "vi") : ?>
                <img src="http://140.238.186.242/signaturegen/imagens/timber-consorcio.png" width="100"></br></br>
            <?php endif; ?>
            <p align="center">

                <?php if ($isa == "1") : ?>
                    <img src="http://140.238.186.242/signaturegen/imagens/sany.png" width="100" height="28"></br></br>
                <?php endif; ?>

                <?php if ($ipo == "1") : ?>
                    <img src="http://140.238.186.242/signaturegen/imagens/ponsse.png" width="90"></br></br>
                <?php endif; ?>

                <?php if ($ifu == "1") : ?>
                    <img src="http://140.238.186.242/signaturegen/imagens/fuchs.png" width="100" height="20"></br></br>
                <?php endif; ?>

                <?php if ($ifl == "1") : ?>
                    <img src="http://140.238.186.242/signaturegen/imagens/timberfleet.png" width="100"></br></br>
                <?php endif; ?>

                <?php if ($ixa == "1") : ?>
                    <img src="http://140.238.186.242/signaturegen/imagens/xag.png" width="80"></br>
                <?php endif; ?>

                <?php if ($eco == "1") : ?>
                    <img src="http://140.238.186.242/signaturegen/imagens/EcoFlow.png" width="80"></br>
                <?php endif; ?>

                <?php if ($pix == "1") : ?>
                    <img src="http://140.238.186.242/signaturegen/imagens/pix4d.png" width="80"></br>
                <?php endif; ?>
            </p>
        </td>
        <td style="border-left: 0.5px solid; padding-left:5px container-fluid">
            <font style="font-family: calibri; color: #000; font-size: 14;">
                <b><?php echo $nome ?></b></br>
                <?php echo $cargo ?></br>
                <b>Telefone:</b> <?php echo $telefone ?>
                <?php if ($celular != "") : ?>|<?php echo $celular ?><?php else : ?><?php endif; ?></br>
                <b>E-mail:</b> <?php echo $email ?></br>
                <b>Skype:</b> <?php echo $skype ?></br>
                <b>EndereÃ§o:</b> <?php echo $endereco1 ?></br><?php echo $endereco2 ?></br>
                <a href="http://www.grupotimber.com.br">www.grupotimber.com.br</a></br>
            </font>
            <font style="font-size: 2;">&nbsp<br></font>
            <a href="https://www.facebook.com/timberforestequipamentos/"><img style="padding-left:2px"
                    src="http://140.238.186.242/signaturegen/imagens/social/facebook.png" width="30" height="30"></a>
            <a href="https://www.linkedin.com/company/timber-forest-equipamentos/?viewAsMember=true"><img
                    style="padding-left:2px" src="http://140.238.186.242/signaturegen/imagens/social/linkedin.png"
                    width="30" height="30"></a>
            <a href="https://www.instagram.com/timber_br/?hl=pt-br"><img style="padding-left:2px"
                    src="http://140.238.186.242/signaturegen/imagens/social/instagram.png" width="30" height="30"></a>
            <a href="https://www.youtube.com/channel/UCflgytZMsL78RUFhECk8IAQ"><img style="padding-left:2px"
                    src="http://140.238.186.242/signaturegen/imagens/social/youtube.png" width="30" height="30"></a>
        </td>
    </tr>
</table>