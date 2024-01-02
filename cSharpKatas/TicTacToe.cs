public class TicTacToe
{
    public TicTacToe()
    {
    }

    internal int IsSolved(int[,] board)
    {
        var i = board.GetLength(0);
        var j = board.GetLength(1);

        var zeroExists = false;
        var vertResult = 0;
        var horResult = 0;

        var firstResult = board[0,0];
        var found = false;

        for (int dim1 = 0; dim1 < i; dim1++)
        {
            firstResult = board[dim1, 0];
            var contSearchHere = true;
            for (var dim2 = 0; dim2 <j; dim2++){
                vertResult = board[dim1, dim2];
                horResult = board[dim2,dim1];
                if (vertResult == 0 || horResult == 0)
                    zeroExists = true;
                if (vertResult == firstResult && vertResult >0 && contSearchHere){
                    found = true;
                }
                else{
                    contSearchHere = false;
                    found = false;
                }
                firstResult = vertResult;
            }
            if (found)
                break;
        }
        if (found)
            return firstResult;

        for (int dim1 = 0; dim1 < j; dim1++)
        {
            firstResult = board[0, dim1];
            var contSearchHere = true;
            for (var dim2 = 0; dim2 <i; dim2++)
            {
                if (vertResult == 0)
                    zeroExists = true;
                vertResult = board[dim2, dim1];
                if (vertResult == firstResult && vertResult >0 && contSearchHere){
                    found = true;
                }
                else{
                    contSearchHere = false;
                    found = false;
                }
                firstResult = vertResult;
            }
            if (found)
                break;
        }
        if (found)
            return firstResult;
        
        firstResult = board[0, 0];
        for (int dim1 = 0; dim1 < j; dim1++)
        {
            vertResult = board[dim1, dim1];
            if (vertResult == firstResult && vertResult >0){
                found = true;
            }
            else{
                found = false;
                break;
            }
            firstResult = vertResult;
        }
        if (found)
            return firstResult;

        firstResult = board[0, 2];    

        for (int dim1 = 0; dim1 < j; dim1++)
        {
            vertResult = board[dim1, j-dim1 - 1];
            if (vertResult == firstResult && vertResult >0){
                found = true;
            }
            else{
                found = false;
                break;
            }
            firstResult = vertResult;
        }
        if (found)
            return firstResult;

        if (zeroExists)
            return -1;

        return 0;
    }
}