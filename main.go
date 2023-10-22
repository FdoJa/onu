package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	pb "github.com/FdoJa/ONU/proto"
	"google.golang.org/grpc"
)

func main() {
	// Dirección OMS
	nameNodeAddr := "10.6.46.108:8080"

	for {
		var status string
		log.Println("Ingresa el estado de las personas que requieres (Infectados/Muertos):")
		_, err := fmt.Scanln(&status)
		if err != nil {
			fmt.Printf("Error al leer la entrada: %v\n", err)
			return
		}

		if strings.ToUpper(status) == "INFECTADOS" {
			status = "INFECTADO"
		} else if strings.ToUpper(status) == "MUERTOS" {
			status = "MUERTO"
		}

		if status == "infectado" || status == "muerto" {
			conn, err := grpc.Dial(nameNodeAddr, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("No se pudo conectar al DataNode: %v", err)
			}
			defer conn.Close()

			client := pb.NewNameNodeClient(conn)

			res, err := client.ConsultarNombres(context.Background(), &pb.Estado_Persona{
				Estado: status,
			})

			if err != nil {
				log.Fatalf("Error al pedir datos en la OMS: %v", err)
			} else {
				fmt.Printf("Recibe algo que no es un error")

				for _, dato := range res.Datos {
					fmt.Printf("Entro al for")
					fmt.Printf("%s %s\n", dato.Nombre, dato.Apellido)
				}
			}

		} else {
			log.Printf("Error: Elige una opción válida")
		}
	}
}
