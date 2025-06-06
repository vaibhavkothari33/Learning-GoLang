package gpt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vaibhavkothari33/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB connection string - contains username, password, and cluster details
const connectionString = "mongodb+srv://vaibhavkothari50:iZqTxT76IgSuoh9k@gobackend.jbas9pp.mongodb.net/?retryWrites=true&w=majority"

// Database name in MongoDB
const dbName = "netflix"

// Collection name (like a table in SQL databases)
const colName = "watchlist"

// Global variable to hold MongoDB collection instance
// This will be used by all functions to interact with the database
var collection *mongo.Collection

// init() function runs automatically when the package is imported
// It establishes connection to MongoDB and sets up the collection
func init() {
	fmt.Println("🚀 Initializing MongoDB connection...")

	// Create client options with the connection string
	clientOption := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB using the client options
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal("❌ Failed to connect to MongoDB:", err)
	}
	fmt.Println("✅ MongoDB connection established successfully!")

	// Get reference to the specific database and collection
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("📦 Collection instance ready - Database:", dbName, "Collection:", colName)
}

// ========== HELPER FUNCTIONS FOR DATABASE OPERATIONS ==========

// insertOneMovie adds a single movie to the watchlist collection
func insertOneMovie(movie model.Netflix) {
	fmt.Println("📝 Adding new movie to watchlist:", movie.Movie)

	// Insert the movie document into the collection
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal("❌ Error inserting movie:", err)
	}

	fmt.Printf("✅ Movie '%s' added successfully with ID: %v\n", movie.Movie, inserted.InsertedID)
}

// updateOneRecord marks a movie as watched by updating its status
func updateOneRecord(movieId string) {
	fmt.Println("🔄 Updating movie status to 'watched' for ID:", movieId)

	// Convert string ID to MongoDB ObjectID format
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal("❌ Invalid movie ID format:", err)
	}

	// Create filter to find the movie by ID
	filter := bson.M{"_id": id}

	// Create update operation to set 'watched' field to true
	update := bson.M{"$set": bson.M{"watched": true}}

	// Execute the update operation
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal("❌ Error updating movie:", err)
	}

	fmt.Printf("✅ Successfully updated %d movie(s) - Movie ID: %s is now marked as watched\n",
		result.ModifiedCount, movieId)
}

// deleteOneMovie removes a single movie from the watchlist
func deleteOneMovie(movieId string) {
	fmt.Println("🗑️ Deleting movie with ID:", movieId)

	// Convert string ID to MongoDB ObjectID
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal("❌ Invalid movie ID format:", err)
	}

	// Create filter to find the movie by ID
	filter := bson.M{"_id": id}

	// Delete the movie document
	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal("❌ Error deleting movie:", err)
	}

	fmt.Printf("✅ Successfully deleted %d movie(s) - Movie ID: %s removed from watchlist\n",
		deleteResult.DeletedCount, movieId)
}

// deleteAllMovies removes all movies from the watchlist collection
func deleteAllMovies() int64 {
	fmt.Println("🧹 Clearing entire watchlist...")

	// Empty filter means "match all documents"
	filter := bson.D{}

	// Delete all documents matching the filter (which is all documents)
	deleteResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal("❌ Error deleting all movies:", err)
	}

	fmt.Printf("✅ Watchlist cleared! Deleted %d movies from database\n", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// getAllMovies retrieves all movies from the watchlist collection
func getAllMovies() []bson.M {
	fmt.Println("📋 Fetching all movies from watchlist...")

	// Empty filter means "get all documents"
	filter := bson.D{}

	// Find all documents in the collection
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal("❌ Error fetching movies:", err)
	}

	// Slice to store all movie documents
	var movies []bson.M

	// Iterate through all documents returned by the query
	for cursor.Next(context.Background()) {
		var movie bson.M

		// Decode each document into a map
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal("❌ Error decoding movie document:", err)
		}

		// Add the movie to our slice
		movies = append(movies, movie)
	}

	// Always close the cursor to free up resources
	defer cursor.Close(context.Background())

	fmt.Printf("✅ Retrieved %d movies from watchlist\n", len(movies))
	return movies
}

// ========== HTTP HANDLER FUNCTIONS (API ENDPOINTS) ==========

// GetMyAllMovies handles GET requests to retrieve all movies
// Endpoint: GET /movies
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🌐 API Call: GET /movies - Retrieving all movies")

	// Set response headers
	w.Header().Set("Content-Type", "application/json") // Fixed content type

	// Get all movies from database
	allMovies := getAllMovies()

	// Convert movies to JSON and send as response
	json.NewEncoder(w).Encode(allMovies)

	fmt.Printf("📤 Response sent: %d movies returned to client\n", len(allMovies))
}

// CreateMovie handles POST requests to add a new movie
// Endpoint: POST /movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🌐 API Call: POST /movie - Adding new movie")

	// Set response headers for JSON and CORS (Cross-Origin Resource Sharing)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")           // Allow requests from any origin
	w.Header().Set("Access-Control-Allow-Methods", "POST")       // Allow POST method

	// Create variable to hold the movie data from request
	var movie model.Netflix

	// Decode JSON from request body into movie struct
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		fmt.Println("❌ Error decoding request body:", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Insert the movie into database
	insertOneMovie(movie)

	// Send the created movie back as confirmation
	json.NewEncoder(w).Encode(movie)

	fmt.Printf("📤 Response sent: Movie '%s' creation confirmed\n", movie.Movie)
}

// MarkAsWatched handles PUT requests to mark a movie as watched
// Endpoint: PUT /movie/{id}
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🌐 API Call: PUT /movie/{id} - Marking movie as watched")

	// Set response headers
	w.Header().Set("Content-Type", "application/json") // Fixed content type
	w.Header().Set("Access-Control-Allow-Methods", "PUT") // Fixed header name

	// Extract URL parameters (the {id} from the route)
	params := mux.Vars(r)
	movieId := params["id"]

	fmt.Printf("🎯 Target movie ID: %s\n", movieId)

	// Update the movie's watched status
	updateOneRecord(movieId)

	// Send back the movie ID as confirmation
	json.NewEncoder(w).Encode(map[string]string{
		"updated_movie_id": movieId,
		"status": "marked_as_watched",
	})

	fmt.Printf("📤 Response sent: Movie %s status update confirmed\n", movieId)
}

// DeleteAMovie handles DELETE requests to remove a specific movie
// Endpoint: DELETE /movie/{id}
func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🌐 API Call: DELETE /movie/{id} - Deleting specific movie")

	// Set response headers
	w.Header().Set("Content-Type", "application/json") // Fixed content type
	w.Header().Set("Access-Control-Allow-Methods", "DELETE") // Fixed header name

	// Extract the movie ID from URL parameters
	params := mux.Vars(r)
	movieId := params["id"]

	fmt.Printf("🎯 Target movie ID for deletion: %s\n", movieId)

	// Delete the movie from database
	deleteOneMovie(movieId)

	// Send confirmation response
	json.NewEncoder(w).Encode(map[string]string{
		"deleted_movie_id": movieId,
		"status": "successfully_deleted",
	})

	fmt.Printf("📤 Response sent: Movie %s deletion confirmed\n", movieId)
}

// DeleteAllMovies handles DELETE requests to clear the entire watchlist
// Endpoint: DELETE /movies
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🌐 API Call: DELETE /movies - Clearing entire watchlist")

	// Set response headers
	w.Header().Set("Content-Type", "application/json") // Fixed content type
	w.Header().Set("Access-Control-Allow-Methods", "DELETE") // Fixed header name

	// Delete all movies and get the count
	count := deleteAllMovies()

	// Send response with deletion count
	json.NewEncoder(w).Encode(map[string]interface{}{
		"deleted_count": count,
		"status": "all_movies_deleted",
		"message": fmt.Sprintf("Successfully deleted %d movies from watchlist", count),
	})

	fmt.Printf("📤 Response sent: Confirmed deletion of %d movies\n", count)
}
