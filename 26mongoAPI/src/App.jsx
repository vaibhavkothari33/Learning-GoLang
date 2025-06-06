import React, { useState, useEffect } from 'react';

const API_BASE = 'http://localhost:4000/api';

const NetflixWatchlist = () => {
  const [movies, setMovies] = useState([]);
  const [newMovie, setNewMovie] = useState({ movie: '', watched: false });
  const [loading, setLoading] = useState(false);

  const fetchMovies = async () => {
    setLoading(true);
    try {
      const res = await fetch(`${API_BASE}/movies`);
      const data = await res.json();
      setMovies(data || []);
    } catch (err) {
      alert('⚠️ Failed to fetch movies.');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const addMovie = async (e) => {
    e.preventDefault();
    if (!newMovie.movie.trim()) return;

    setLoading(true);
    try {
      await fetch(`${API_BASE}/movies`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(newMovie),
      });
      setNewMovie({ movie: '', watched: false });
      fetchMovies();
    } catch (err) {
      alert('⚠️ Failed to add movie.');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const markAsWatched = async (id) => {
    setLoading(true);
    try {
      await fetch(`${API_BASE}/movies/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
      });
      fetchMovies();
    } catch (err) {
      alert('⚠️ Failed to mark as watched.');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const deleteMovie = async (id) => {
    if (!window.confirm('Delete this movie?')) return;

    setLoading(true);
    try {
      await fetch(`${API_BASE}/movies/${id}`, { method: 'DELETE' });
      fetchMovies();
    } catch (err) {
      alert('⚠️ Failed to delete movie.');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const deleteAllMovies = async () => {
    if (!window.confirm('Delete ALL movies?')) return;

    setLoading(true);
    try {
      await fetch(`${API_BASE}/movies`, { method: 'DELETE' });
      fetchMovies();
    } catch (err) {
      alert('⚠️ Failed to delete all movies.');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchMovies();
  }, []);

  return (
    <div style={styles.container}>
      <h1 style={styles.heading}>🎬 Netflix Watchlist</h1>

      {/* Add Movie */}
      <form onSubmit={addMovie} style={styles.form}>
        <input
          type="text"
          placeholder="Enter movie name"
          value={newMovie.movie}
          onChange={(e) => setNewMovie({ ...newMovie, movie: e.target.value })}
          style={styles.input}
        />
        <label style={styles.checkboxLabel}>
          <input
            type="checkbox"
            checked={newMovie.watched}
            onChange={(e) => setNewMovie({ ...newMovie, watched: e.target.checked })}
          />
          Watched
        </label>
        <button type="submit" disabled={loading} style={styles.addButton}>
          ➕ Add
        </button>
      </form>

      {/* Controls */}
      <div style={styles.controls}>
        <button onClick={fetchMovies} disabled={loading} style={styles.button}>
          🔄 Refresh
        </button>
        <button onClick={deleteAllMovies} disabled={loading} style={styles.dangerButton}>
          🗑️ Delete All
        </button>
      </div>

      {/* Movie List */}
      <h3 style={{ marginTop: 30 }}>Your Watchlist ({movies.length})</h3>
      <div style={styles.grid}>
        {loading && <p>Loading...</p>}
        {!loading && movies.length === 0 && <p>No movies yet. Add some!</p>}

        {movies.map((movie) => (
          <div key={movie._id} style={{ ...styles.card, backgroundColor: movie.watched ? '#000000' :"#000000" }}>
            <h4>{movie.movie}</h4>
            <p>Status: <strong>{movie.watched ? '✅ Watched' : '⏳ Not Watched'}</strong></p>
            <div style={styles.cardButtons}>
              <button
                onClick={() => markAsWatched(movie._id)}
                disabled={loading || movie.watched}
                style={styles.smallButton}
              >
                ✅ Watch
              </button>
              <button
                onClick={() => deleteMovie(movie._id)}
                disabled={loading}
                style={styles.deleteButton}
              >
                ❌ Delete
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

const styles = {
  container: {
    padding: 30,
    fontFamily: 'Arial, sans-serif',
    maxWidth: '900px',
    margin: '0 auto',
  },
  heading: {
    fontSize: '2.2rem',
    marginBottom: 20,
  },
  form: {
    display: 'flex',
    gap: 10,
    flexWrap: 'wrap',
    marginBottom: 20,
  },
  input: {
    padding: 10,
    flexGrow: 1,
    minWidth: '200px',
    borderRadius: 4,
    border: '1px solid #ccc',
  },
  checkboxLabel: {
    display: 'flex',
    alignItems: 'center',
    gap: 5,
  },
  addButton: {
    backgroundColor: '#007bff',
    color: '#fff',
    padding: '10px 16px',
    border: 'none',
    borderRadius: 4,
    cursor: 'pointer',
  },
  controls: {
    marginBottom: 20,
    display: 'flex',
    gap: 10,
  },
  button: {
    padding: '10px 16px',
    backgroundColor: '#4caf50',
    color: 'white',
    border: 'none',
    borderRadius: 4,
    cursor: 'pointer',
  },
  dangerButton: {
    padding: '10px 16px',
    backgroundColor: '#f44336',
    color: 'white',
    border: 'none',
    borderRadius: 4,
    cursor: 'pointer',
  },
  grid: {
    display: 'grid',
    gap: 20,
    gridTemplateColumns: 'repeat(auto-fit, minmax(240px, 1fr))',
  },
  card: {
    border: '1px solid #ddd',
    borderRadius: 8,
    padding: 16,
    boxShadow: '0 2px 6px rgba(0,0,0,0.1)',
  },
  cardButtons: {
    display: 'flex',
    gap: 10,
    marginTop: 10,
  },
  smallButton: {
    padding: '6px 12px',
    backgroundColor: '#2196f3',
    color: 'white',
    border: 'none',
    borderRadius: 4,
    cursor: 'pointer',
  },
  deleteButton: {
    padding: '6px 12px',
    backgroundColor: '#e53935',
    color: 'white',
    border: 'none',
    borderRadius: 4,
    cursor: 'pointer',
  },
};

export default NetflixWatchlist;
