package com.example.upcycleconnectpro

import android.content.Context
import android.graphics.Color
import android.graphics.Paint
import android.graphics.pdf.PdfDocument
import androidx.core.content.FileProvider
import com.example.upcycleconnectpro.network.Invoice
import java.io.File
import java.io.FileOutputStream

// Génère une facture PDF en local (sans aucune dépendance externe) grâce à PdfDocument du SDK Android.
object InvoicePdfGenerator {

    // Dimensions d'une page A4 en points (1/72 de pouce)
    private const val PAGE_WIDTH = 595
    private const val PAGE_HEIGHT = 842

    // Génère le fichier PDF et renvoie le File créé
    fun generate(context: Context, invoice: Invoice, nomClient: String): File {
        val document = PdfDocument()
        val pageInfo = PdfDocument.PageInfo.Builder(PAGE_WIDTH, PAGE_HEIGHT, 1).create()
        val page = document.startPage(pageInfo)
        val canvas = page.canvas

        val titrePaint = Paint().apply {
            color = Color.parseColor("#2E7D32")
            textSize = 26f
            isFakeBoldText = true
        }
        val labelPaint = Paint().apply {
            color = Color.parseColor("#888888")
            textSize = 12f
        }
        val textePaint = Paint().apply {
            color = Color.BLACK
            textSize = 14f
        }
        val grasPaint = Paint().apply {
            color = Color.BLACK
            textSize = 16f
            isFakeBoldText = true
        }
        val lignePaint = Paint().apply {
            color = Color.parseColor("#DDDDDD")
            strokeWidth = 1f
        }

        val marge = 50f
        var y = 70f

        // En-tête
        canvas.drawText("UpcycleConnect Pro", marge, y, titrePaint)
        y += 24f
        canvas.drawText("Facture de récupération", marge, y, labelPaint)

        // Numéro + date (à droite)
        canvas.drawText("Facture n° ${invoice.id}", PAGE_WIDTH - 200f, 70f, textePaint)
        canvas.drawText("Date : ${invoice.date}", PAGE_WIDTH - 200f, 92f, textePaint)

        y += 30f
        canvas.drawLine(marge, y, PAGE_WIDTH - marge, y, lignePaint)

        // Client
        y += 40f
        canvas.drawText("CLIENT", marge, y, labelPaint)
        y += 22f
        canvas.drawText(nomClient, marge, y, textePaint)

        // Détail
        y += 50f
        canvas.drawText("DÉTAIL", marge, y, labelPaint)
        y += 26f
        canvas.drawText("Article", marge, y, grasPaint)
        canvas.drawText("Montant", PAGE_WIDTH - 150f, y, grasPaint)
        y += 10f
        canvas.drawLine(marge, y, PAGE_WIDTH - marge, y, lignePaint)

        y += 28f
        canvas.drawText(invoice.titre, marge, y, textePaint)
        canvas.drawText(String.format("%.2f €", invoice.montant), PAGE_WIDTH - 150f, y, textePaint)

        y += 24f
        canvas.drawText("Dont commission plateforme (5%)", marge, y, labelPaint)
        canvas.drawText(String.format("%.2f €", invoice.commission), PAGE_WIDTH - 150f, y, labelPaint)

        // Total
        y += 30f
        canvas.drawLine(marge, y, PAGE_WIDTH - marge, y, lignePaint)
        y += 30f
        canvas.drawText("TOTAL PAYÉ", marge, y, grasPaint)
        canvas.drawText(String.format("%.2f €", invoice.montant), PAGE_WIDTH - 150f, y, grasPaint)

        // Pied de page
        canvas.drawText(
            "Merci de votre confiance — UpcycleConnect",
            marge,
            PAGE_HEIGHT - 50f,
            labelPaint
        )

        document.finishPage(page)

        // Sauvegarde dans le dossier privé de l'app (pas besoin de permission)
        val file = File(context.getExternalFilesDir(null), "facture_${invoice.id}.pdf")
        document.writeTo(FileOutputStream(file))
        document.close()

        return file
    }

    // Renvoie une URI partageable (via FileProvider) pour ouvrir le PDF dans une autre app
    fun getShareableUri(context: Context, file: File) =
        FileProvider.getUriForFile(context, "${context.packageName}.provider", file)
}
