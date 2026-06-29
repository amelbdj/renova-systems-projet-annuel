package admin

import (
	"bytes"
	"fmt"
	"os"
	"time"
	"upcycleconnect/bdd"

	"github.com/go-pdf/fpdf"
)

func GenerateInvoicePDF(orderID int, buyerID int, articleTitre string, montant float64) (string, error) {
	var nom, prenom string
	bdd.Db.QueryRow("SELECT COALESCE(nom, ''), COALESCE(prenom, '') FROM utilisateur WHERE id = ?", buyerID).Scan(&nom, &prenom)

	pdf := fpdf.New("P", "mm", "A4", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(0, 12, "ReNova Systems")
	pdf.Ln(14)

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, tr("Facture"))
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 11)
	pdf.Cell(0, 7, tr(fmt.Sprintf("Facture N : FAC-%d", orderID)))
	pdf.Ln(7)
	pdf.Cell(0, 7, tr("Date : "+time.Now().Format("02/01/2006")))
	pdf.Ln(7)
	pdf.Cell(0, 7, tr("Client : "+prenom+" "+nom))
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(130, 8, tr("Article"), "1", 0, "L", false, 0, "")
	pdf.CellFormat(40, 8, tr("Montant"), "1", 1, "R", false, 0, "")

	pdf.SetFont("Arial", "", 11)
	pdf.CellFormat(130, 8, tr(articleTitre), "1", 0, "L", false, 0, "")
	pdf.CellFormat(40, 8, fmt.Sprintf("%.2f EUR", montant), "1", 1, "R", false, 0, "")

	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(130, 8, tr("Total"), "1", 0, "R", false, 0, "")
	pdf.CellFormat(40, 8, fmt.Sprintf("%.2f EUR", montant), "1", 1, "R", false, 0, "")
	pdf.Ln(15)

	pdf.SetFont("Arial", "I", 9)
	pdf.MultiCell(0, 5, tr("Merci d'avoir utilise ReNova Systems."), "", "L", false)

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return "", err
	}

	if err := os.MkdirAll("documents/doubles", 0755); err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("facture_%d_%d.pdf", orderID, time.Now().Unix())
	original := "documents/" + fileName
	double := "documents/doubles/" + fileName

	if err := os.WriteFile(original, buf.Bytes(), 0644); err != nil {
		return "", err
	}
	if err := os.WriteFile(double, buf.Bytes(), 0644); err != nil {
		return "", err
	}

	urlPdf := "/view-documents/doubles/" + fileName
	bdd.Db.Exec("INSERT INTO document (id_user, type_doc, url_pdf, id_commande) VALUES (?, ?, ?, ?)", buyerID, "facture", urlPdf, orderID)

	return urlPdf, nil
}

func GenerateContractPDF(userID int, plan string, abonnementID int) (string, error) {
	var nom, prenom string
	bdd.Db.QueryRow("SELECT COALESCE(nom, ''), COALESCE(prenom, '') FROM utilisateur WHERE id = ?", userID).Scan(&nom, &prenom)

	planLabel := "Premium"
	prix := 25.0
	if plan == "plus" {
		planLabel = "Plus"
		prix = 45.0
	} else if plan == "pro" {
		planLabel = "Pro"
		prix = 99.0
	}

	pdf := fpdf.New("P", "mm", "A4", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(0, 12, "ReNova Systems")
	pdf.Ln(14)

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, tr("Contrat d'abonnement"))
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 11)
	pdf.Cell(0, 7, tr("Date : "+time.Now().Format("02/01/2006")))
	pdf.Ln(7)
	pdf.Cell(0, 7, tr("Client : "+prenom+" "+nom))
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(130, 8, tr("Abonnement"), "1", 0, "L", false, 0, "")
	pdf.CellFormat(40, 8, tr("Prix / mois"), "1", 1, "R", false, 0, "")

	pdf.SetFont("Arial", "", 11)
	pdf.CellFormat(130, 8, tr("Abonnement "+planLabel), "1", 0, "L", false, 0, "")
	pdf.CellFormat(40, 8, fmt.Sprintf("%.2f EUR", prix), "1", 1, "R", false, 0, "")
	pdf.Ln(15)

	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(0, 6, tr("Cet abonnement est sans engagement et peut etre resilie a tout moment depuis votre espace professionnel. Le montant est preleve chaque mois via Stripe."), "", "L", false)
	pdf.Ln(6)

	pdf.SetFont("Arial", "I", 9)
	pdf.MultiCell(0, 5, tr("Merci de votre confiance - ReNova Systems."), "", "L", false)

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return "", err
	}

	if err := os.MkdirAll("documents/doubles", 0755); err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("contrat_%d_%d.pdf", userID, time.Now().Unix())
	original := "documents/" + fileName
	double := "documents/doubles/" + fileName

	if err := os.WriteFile(original, buf.Bytes(), 0644); err != nil {
		return "", err
	}
	if err := os.WriteFile(double, buf.Bytes(), 0644); err != nil {
		return "", err
	}

	urlPdf := "/view-documents/doubles/" + fileName
	bdd.Db.Exec("INSERT INTO document (id_user, type_doc, url_pdf, id_commande) VALUES (?, ?, ?, ?)", userID, "contrat", urlPdf, abonnementID)

	return urlPdf, nil
}
