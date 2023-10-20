package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	pb "github.com/FdoJa/ONU/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("Fallo en escuchar: %v", err)
	}

	s := grpc.NewServer()

	fmt.Println("Servidor ONU escuchando en :80")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fallo en serve: %v", err)
	}

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

		if strings.ToUpper(status) == "INFECTADOS" || strings.ToUpper(status) == "MUERTOS" {
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
				for _, dato := range res.Datos {
					fmt.Printf("%s %s\n", dato.Nombre, dato.Apellido)
				}
			}

		} else {
			log.Printf("Error: Elige una opción válida")
		}
	}
}
