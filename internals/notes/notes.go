package notes

import (
	"fmt"
	"strings"

	"github.com/romaingallez/scform_api/internals/models"
	"github.com/romaingallez/scform_api/internals/utils"
)

func GetNotes(scPage models.Page) (Matieres []models.Matiere, err error) {

	page := scPage.Page
	page.MustWaitDOMStable()

	page.MustEval(`() => GoTo('Eleve/MesNotes.aspx')`)
	page.MustWaitLoad()

	page.MustEval(`
	() => {

	const tdElement = Array.from(document.querySelectorAll('#MainContent_RadioButtonAffichage td')).find(td => td.textContent.trim() === 'par matière');
	if (tdElement) {
		const inputElement = tdElement.querySelector('input');
		if (inputElement) {
	    	inputElement.click();
		}
	}
	}
	`)

	page.MustWaitLoad()

	TableNotes := page.MustElements("table.AfficheInfoEnMieux")

	for i, tableNotes := range TableNotes {
		var Matiere models.Matiere

		matiereSelector := fmt.Sprintf(`#MainContent_ListViewLstMatiere_NomCompletLabel_%d`, i)
		MatiereTitle := tableNotes.MustElement(matiereSelector).MustText()
		Matiere.NomMatiere = strings.Join(strings.Split(MatiereTitle, "\n"), " ")
		divNotes := tableNotes.MustElements("#DivNOTE")
		for _, divNotes := range divNotes {
			var Note models.Note
			// document.querySelector("#MainContent_ListViewLstMatiere_ListViewLstNoteParDevoir_0_ctrl0_0_Label1_0")
			// getElement :=
			// selectNote := fmt.Sprintf(`() => getElement(%d, %d)`, i, i)

			// eval getElement then use selectNote to call it
			note := divNotes.MustEval(fmt.Sprintf(`
			() => {
				function getNoteText(index1, index2) {
					const selector = `+"`#MainContent_ListViewLstMatiere_ListViewLstNoteParDevoir_${index1}_ctrl0_${index2}_Label1_0`"+`;
					const element = document.querySelector(selector);
					if (element) {
						return element.textContent;
					} else {
						return -1;
					}
				}
				return getNoteText(%d, %d);
			}
		`, i, i)).Str()

			note = strings.Join(strings.Split(note, "\n"), " ")

			noteMax := divNotes.MustEval(fmt.Sprintf(`
			() => {
			function getNoteText(index1, index2) {
				const selector = `+"`#MainContent_ListViewLstMatiere_ListViewLstNoteParDevoir_${index1}_ctrl0_${index2}_Label2_0`"+`;
				const element = document.querySelector(selector);
				if (element) {
					return element.textContent;
				} else {
					return -1;
				}
			}
			return getNoteText(%d, %d);
		}
		`, i, i)).Str()

			Note.Note = note

			// remote `/` from noteMax
			Note.NoteMax = strings.ReplaceAll(noteMax, "/", "")

			Note.NoteMax = strings.Join(strings.Split(Note.NoteMax, "\n"), " ")

			Coeef := divNotes.MustEval(fmt.Sprintf(`
			() => {
				function getNoteText(index1, index2) {
					const selector = `+"`#MainContent_ListViewLstMatiere_ListViewLstNoteParDevoir_${index1}_ctrl0_${index2}_Label3_0`"+`;
					const element = document.querySelector(selector);
					if (element) {
						return element.textContent;
					} else {
						return -1;
					}
				}
				return getNoteText(%d, %d);
			}
			`, i, i)).Str()

			Note.Coeef = strings.ReplaceAll(Coeef, "coeff. ", "")
			// split Note.Coeef on \n trim space then join with space
			Note.Coeef = strings.Join(strings.Split(Note.Coeef, "\n"), " ")

			Desc := divNotes.MustEval(fmt.Sprintf(`
			() => {
				function getNoteText(index1, index2) {
					const selector = `+"`#MainContent_ListViewLstMatiere_ListViewLstNoteParDevoir_${index1}_ctrl0_${index2}_Label9_0`"+`;
					const element = document.querySelector(selector);
					if (element) {
						return element.textContent;
					} else {
						return -1;
					}
				}
				return getNoteText(%d, %d);
			}
				`, i, i)).Str()

			Note.Desc = strings.Join(strings.Split(Desc, "\n"), " ")

			Note.Type = divNotes.MustEval(fmt.Sprintf(`
			() => {
				function getNoteText(index1, index2) {
					const selector = `+"`#MainContent_ListViewLstMatiere_ListViewLstNoteParDevoir_${index1}_ctrl0_${index2}_Label8_0`"+`;
					const element = document.querySelector(selector);
					if (element) {
						return element.textContent;
					} else {
						return -1;
					}
				}
				return getNoteText(%d, %d);
			}
		`, i, i)).Str()

			Note.Type = strings.Join(strings.Split(Note.Type, "\n"), " ")

			Note.Type = utils.RemoveExtraSpaces(Note.Type)

			Note.Name = divNotes.MustEval(fmt.Sprintf(`
			() => {
				function getNoteText(index1, index2) {
					const selector = `+"`#MainContent_ListViewLstMatiere_ListViewLstNoteParDevoir_${index1}_ctrl0_${index2}_Label7_0`"+`;
					const element = document.querySelector(selector);
					if (element) {
						return element.textContent;
					} else {
						return -1;
					}
				}
				return getNoteText(%d, %d);
			}
	`, i, i)).Str()

			Note.Name = strings.Join(strings.Split(Note.Name, "\n"), " ")
			Note.Name = utils.RemoveExtraSpaces(Note.Name)

			Matiere.Notes = append(Matiere.Notes, Note)
		}

		Matieres = append(Matieres, Matiere)

	}

	return Matieres, nil
}
